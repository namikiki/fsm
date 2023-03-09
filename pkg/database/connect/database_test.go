package connect

import (
	"database/sql"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

func CreateLocalTable(sqlConn *sqlx.DB, tabName, path string) {
	createFolderTabSql := fmt.Sprintf("CREATE TABLE %s (id TEXT,fol TEXT,pfol TEXT,step INTEGER,times INTEGER)", tabName+"folder")
	createFileTabSql := fmt.Sprintf("CREATE TABLE %s (id TEXT,fol TEXT,pfol TEXT,step INTEGER,times INTEGER)", tabName+"file")

	insertFolderSql := fmt.Sprintf("INSERT INTO %s (id,fol, pfol,step,times) VALUES (?,?,?,?,?)", tabName+"folder")
	insertFileSql := fmt.Sprintf("INSERT INTO %s (id,fol, pfol,step,times) VALUES (?,?,?,?,?)", tabName+"file")

	if _, err := sqlConn.Exec(createFolderTabSql); err != nil {
		log.Fatal(err)
	}
	if _, err := sqlConn.Exec(createFileTabSql); err != nil {
		log.Fatal(err)
	}

	suff, base := filepath.Split(path)
	_, _ = sqlConn.Exec(insertFolderSql, uuid.New().String(), base, suff, 0, 0)

	if err := WalkDir(path, func(path string, d fs.FileInfo, err error) error {
		split := strings.Split(path[len(suff):], "\\")
		splen := len(split)

		if d.IsDir() {
			_, err = sqlConn.Exec(insertFolderSql, uuid.New().String(), split[splen-1], split[splen-2], splen-1, d.ModTime().Unix())
		} else {
			_, err = sqlConn.Exec(insertFileSql, uuid.New().String(), split[splen-1], split[splen-2], splen-1, d.ModTime().Unix())
		}

		return err
	}); err != nil {
		log.Println(err)
		return
	}

}

func CreateLocalItemTable(sqlConn *sqlx.DB, tabName, path string) {
	createTabSql := fmt.Sprintf("CREATE TABLE %s (id TEXT,is_folder INTEGER,name TEXT,parent TEXT,step INTEGER,times INTEGER)", tabName)
	insertSql := fmt.Sprintf("INSERT INTO %s (id,is_folder,name,parent,step,times) VALUES (?,?,?,?,?,?)", tabName)

	if _, err := sqlConn.Exec(createTabSql); err != nil {
		log.Fatal(err)
	}

	suff, base := filepath.Split(path)
	_, _ = sqlConn.Exec(insertSql, uuid.New().String(), 1, base, suff, 0, 0)

	if err := WalkDir(path, func(path string, d fs.FileInfo, err error) error {
		split := strings.Split(path[len(suff):], "\\")
		splen := len(split)

		if d.IsDir() {
			_, err = sqlConn.Exec(insertSql, uuid.New().String(), 1, split[splen-1], split[splen-2], splen-1, d.ModTime().Unix())
		} else {
			_, err = sqlConn.Exec(insertSql, uuid.New().String(), 0, split[splen-1], split[splen-2], splen-1, d.ModTime().Unix())
		}

		return err
	}); err != nil {
		log.Println(err)
		return
	}

}

type Mem struct {
	Id    string
	Fol   string
	Pfol  string
	Step  uint
	Times int64
	Res   int
}

type FileChange struct {
	ID           string
	State        string
	Folder       string
	ParentFolder string
	Step         uint
	Times        int64
}

func TestTT(t *testing.T) {

	path := "C:\\Users\\surflabom\\Desktop\\f1"
	locTempFilename := "file:ass.db?mode=rwc&cache=shared&_fk=1&_cache_size=2000"

	sotDb, err := sqlx.Connect("sqlite3", locTempFilename)
	if err != nil {
		t.Fatal("Failed to open the destination database:", err)
	}
	CreateLocalItemTable(sotDb, "old", path)
	CreateLocalItemTable(sotDb, "new", path)
	//
	//return

	//loc copy to mem
	memTempFilename := "file:acc.db?mode=memory&cache=shared&_fk=1&_cache_size=2000"
	//destDb, err := sqlx.Connect("sqlite3", memTempFilename)
	//if err != nil {
	//	t.Fatal("Failed to open the destination database:", err)
	//}

	return

	_, mem := GetDBConn(locTempFilename, memTempFilename)

	//rows, _ := mem.Query("select * from oldtab;")
	//
	//var res []Mem
	//for rows.Next() {
	//	var s Mem
	//	if err := rows.Scan(&s.Id, &s.Fol, &s.Pfol, &s.Step); err != nil {
	//		log.Fatal(err)
	//	}
	//	res = append(res, s)
	//}
	//log.Println(res)

	//scan now
	CreateLocalTable(mem, "new", path)

	//rows, _ := mem.Query("select * from new;")
	//var res []Mem
	//for rows.Next() {
	//	var s Mem
	//	if err := rows.Scan(&s.Id, &s.Fol, &s.Pfol, &s.Step); err != nil {
	//		log.Fatal(err)
	//	}
	//	res = append(res, s)
	//}
	//log.Println(res)

	file := map[string]*FileChange{}

	//diff
	rows, err := mem.Queryx("select fol,pfol,step,times,res from ( select fol,pfol,step,times, 1 as res from oldfile union all select fol,pfol,step,times,-1 as res from newfile) group by fol,pfol,step,times having sum(res);")
	if err != nil {
		log.Printf("%v", err)
	}

	log.Println("111----1111")
	for rows.Next() {
		var f Mem
		rows.StructScan(&f)

		sta := "add"
		if f.Res > 0 {
			sta = "update"
		}
		file[f.Id] = &FileChange{
			ID:           f.Id,
			State:        sta,
			Folder:       f.Fol,
			ParentFolder: f.Pfol,
			Step:         f.Step,
			Times:        f.Times,
		}
		log.Println(f)
	}

	log.Println("111----2222")
	rows, err = mem.Queryx("select id,fol,pfol,step,res from ( select id,fol,pfol,step, 1 as res from oldfile union all select id,fol,pfol,step,-1 as res from newfile) group by fol,pfol,step having sum(res);")
	if err != nil {
		log.Printf("%v", err)
	}

	for rows.Next() {
		var f Mem
		rows.StructScan(&f)
		if f.Res > 0 {
			file[f.Id].State = "delete"
		}
		log.Println(f)
	}

	for id, change := range file {
		log.Println(id, change.Folder, change.State)
	}

	//rows, _ := mem.Query("select id,fol,pfol,step,res from ( select id,fol,pfol,step,1 as res from oldfolder union all select id,fol,pfol,step,-1 as res from newfolder) group by fol,pfol,step having sum(res);")

}

type Ttbale struct {
	Id    int
	Value string
	Has   int
}

func TestKS(t *testing.T) {

	WalkDir("C:\\Users\\surflabom\\Desktop\\f1", func(path string, d fs.FileInfo, err error) error {
		log.Println(path, d.IsDir())
		return err
	})
}

var SkipDir error = fs.SkipDir

type WalkDirFunc func(path string, d fs.FileInfo, err error) error

func WalkDir(root string, fn WalkDirFunc) error {
	info, err := os.Lstat(root)
	if err != nil {
		return err
	}
	return walkDir(root, info, fn)
}

func walkDir(path string, d fs.FileInfo, walkDirFn WalkDirFunc) error {

	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		// Second call, to report ReadDir error.
		err = walkDirFn(path, d, err)
		if err != nil {
			if err == SkipDir && d.IsDir() {
				err = nil
			}
			return err
		}
	}

	for _, d1 := range dirs {
		path1 := filepath.Join(path, d1.Name())

		if err = walkDirFn(path1, d1, err); err != nil {
			return err
		}

		if !d1.IsDir() {
			continue
		}

		if err := walkDir(path1, d1, walkDirFn); err != nil {
			if err == SkipDir {
				break
			}
			return err
		}
	}
	return err
}

//func walkDir(path string, d fs.FileInfo, walkDirFn WalkDirFunc) error {
//
//	dirs, err := ioutil.ReadDir(path)
//	if err != nil {
//		// Second call, to report ReadDir error.
//		//err = walkDirFn(path, d, err)
//		//if err != nil {
//		//	if err == SkipDir && d.IsDir() {
//		//		err = nil
//		//	}
//		//	return err
//		//}
//		return err
//	}
//
//	for _, dn := range dirs {
//		nextPath := filepath.Join(path, dn.Name())
//
//		if err = walkDirFn(path, d, err); err != nil {
//			return err
//		}
//
//		if !d.IsDir() {
//			continue
//		}
//
//		if err := walkDir(nextPath, dn, walkDirFn); err != nil {
//			if err == SkipDir {
//				break
//			}
//			return err
//		}
//
//		//if err = walkDirFn(nextPath, d, err); err != nil {
//		//	if err == SkipDir && d.IsDir() {
//		//		err = nil
//		//	}
//		//	return err
//		//}
//		//
//		//if !d.IsDir() {
//		//	con
//		//}
//		//
//		//if err := walkDir(nextPath, dn, walkDirFn); err != nil {
//		//	if err == SkipDir {
//		//		break
//		//	}
//		//	return err
//		//}
//	}
//
//	return err
//}

//func SuffBase(filepaths string) (string, string) {
//	base := filepath.Base(filepaths)
//	suff := filepaths[0 : len(filepaths)-len(base)]
//	return suff, base
//}

// 遍历路径至数据库
func WalkDirToDataBase(sqlConn *sqlx.DB, Path, tab string) error {
	createTableSql := fmt.Sprintf("CREATE TABLE %s (id TEXT,fol TEXT,pfol TEXT,step INTEGER,isdir INTEGER,md5 TEXT,modtime INTEGER)", tab)
	insertTableSql := fmt.Sprintf("INSERT INTO %s (id,fol, pfol,step,isdir,md5,modtime) VALUES (?,?,?,?,?,?,?)", tab)

	p, _ := os.Stat(Path)

	if _, err := sqlConn.Exec(createTableSql); err != nil {
		log.Fatal(err)
	}

	dir, file := filepath.Split(Path)
	if _, err := sqlConn.Exec(insertTableSql, "", file, dir, 0, 1, "", p.ModTime().Unix()); err != nil {
		log.Fatal(err)
	}
	dirlen := len(dir)
	pathSeparator := string(os.PathSeparator)

	err := WalkDir(Path, func(path string, d fs.FileInfo, err error) error {
		split := strings.Split(path[dirlen:], pathSeparator)
		splen := len(split)

		if d.IsDir() {
			_, err = sqlConn.Exec(insertTableSql, "", split[splen-1], split[splen-2], splen-1, 1, "", d.ModTime().Unix())
		} else {
			_, err = sqlConn.Exec(insertTableSql, "", split[splen-1], split[splen-2], splen-1, 0, "", d.ModTime().Unix())
		}

		//time.Sleep(time.Second)
		//log.Printf(path, split[splen-2], split[splen-1], splen-1)

		return err
	})

	return err
}

//// func A(sqlConn *sql.DB, tab1, tab2 string) {
//func A(sqlConn *sql.DB, tab1, tab2 string) {
//	filepaths := "/Users/zylzyl/Desktop/markdown"
//
//	tab1Sql := fmt.Sprintf("CREATE TABLE %s (id TEXT,fol TEXT,pfol TEXT,step INTEGER)", tab1)
//	tab2Sql := fmt.Sprintf("CREATE TABLE %s (id TEXT PRIMARY KEY,fol TEXT,pfol TEXT,step INTEGER)", tab2)
//
//	_, err := sqlConn.Exec(tab1Sql)
//	if err != nil {
//		log.Fatal(err)
//	}
//	suff, base := filepath.Split(filepaths)
//
//	log.Printf("suff %s,base %s\n", suff, base)
//	_, err = sqlConn.Exec("INSERT INTO test (id,fol, pfol,step) VALUES (?,?,?,?)", "", base, suff, 0)
//	if err := WalkDir(filepaths, func(path string, d fs.FileInfo, err error) error {
//		split := strings.Split(path[len(suff):], "/")
//		splen := len(split)
//
//		_, err = sqlConn.Exec("INSERT INTO test (id,fol, pfol,step) VALUES (?,?,?,?)", "", split[splen-2], split[splen-1], splen-1)
//		return err
//	}); err != nil {
//		log.Println(err)
//		return
//	}
//
//	filepath2 := "/Users/zylzyl/Desktop/markdown"
//	_, err = sqlConn.Exec(tab2Sql)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	suff2, base2 := filepath.Split(filepath2)
//
//	log.Printf("suff %s,base %s\n", suff2, base2)
//
//	_, err = sqlConn.Exec("INSERT INTO test2 (id,fol, pfol,step) VALUES (?,?,?,?)", uuid.New().String(), base2, suff2, 0)
//	if err := WalkDir(filepath2, func(path string, d fs.FileInfo, err error) error {
//		split := strings.Split(path[len(suff2):], "/")
//		splen := len(split)
//
//		_, err = sqlConn.Exec("INSERT INTO test2 (id,fol, pfol,step) VALUES (?,?, ?,?)", uuid.New().String(), split[splen-2], split[splen-1], splen-1)
//		if err != nil {
//			log.Println(err)
//		}
//		return err
//	}); err != nil {
//		log.Println(err)
//		return
//	}
//
//	rows, err := sqlConn.Query("select id,fol,pfol,step,res from ( select id,fol,pfol,step, 1 as res from test union all select id,fol,pfol,step,-1 as res from test2) group by fol,pfol,step having sum(res);")
//	//rows, err := sqlConn.Query("select * from test2;")
//
//	var res []Mem
//	for rows.Next() {
//		var s Mem
//		if err := rows.Scan(&s.Id, &s.Fol, &s.Pfol, &s.Step); err != nil {
//			log.Fatal(err)
//		}
//		res = append(res, s)
//	}
//	log.Println(res)
//
//	//_, err = sqlConn.Exec(tab1Sql)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//}

type DirItemMeta struct {
	Id    string
	Fol   string
	Pfol  string
	Step  uint
	Isdir uint
	Md5   string
	Res   int
}

func CompareLocalSyncPathChange(sqlConn *sqlx.DB, tab1, tab2 string) {

	sprintf := fmt.Sprintf(
		"select id,fol,pfol,step,isdir,md5,res from ( select id,fol,pfol,step,isdir,md5, 1 as res from %s union all select id,fol,pfol,step,isdir,md5,-1 as res from %s) group by fol,pfol,step having sum(res);", tab1, tab2)
	rows, err := sqlConn.Query(sprintf)
	if err != nil {
		log.Println("232", err)
	}

	var res []DirItemMeta
	for rows.Next() {
		var s DirItemMeta
		if err := rows.Scan(&s.Id, &s.Fol, &s.Pfol, &s.Step, &s.Isdir, &s.Md5, &s.Res); err != nil {
			log.Fatal(err)
		}
		log.Println(s)
		res = append(res, s)
	}
	log.Println(res)

}

func TestCompareLocalSyncPathChange(t *testing.T) {
	destTempFilename := "file:ass.db?mode=rwc&cache=shared&_fk=1&_cache_size=2000"

	destDb, err := sqlx.Connect("sqlite3", destTempFilename)
	if err != nil {
		t.Fatal("Failed to open the destination database:", err)
	}
	CompareLocalSyncPathChange(destDb, "ts1", "ts2")
}

func TestS(t *testing.T) {

	//destTempFilename := "file:aff.db?mode=memory&cache=shared&_fk=1&_cache_size=2000"
	//var driverName = "sqlite3_with_hook"
	destTempFilename := "file:ass.db?mode=rwc&cache=shared&_fk=1&_cache_size=2000"

	destDb, err := sqlx.Connect("sqlite3", destTempFilename)
	if err != nil {
		t.Fatal("Failed to open the destination database:", err)
	}

	WalkDirToDataBase(destDb, "/Users/zylzyl/Desktop/GolangProjects/Project01/cmd/sync", "ts1")
	//WalkDirToDataBase(destDb, "/Users/zylzyl/Desktop/testone", "ts2")

	//CompareLocalSyncPathChange(destDb, "ts1", "ts2")

}

func GetDBConn(localDataSource, memDataSource string) (*sqlx.DB, *sqlx.DB) {
	var err error
	var driverCons []*sqlite3.SQLiteConn
	var driverName = "sqlite3_with_hook"

	sql.Register(driverName, &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			driverCons = append(driverCons, conn)
			return nil
		}})

	srcDb, err := sqlx.Connect(driverName, localDataSource)
	if err != nil {
		log.Fatal("Failed to open the source database:", err)
	}
	err = srcDb.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the source database:", err)
	}

	destDb, err := sqlx.Connect(driverName, memDataSource)
	if err != nil {
		log.Fatal("Failed to open the destination database:", err)
	}
	err = destDb.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the destination database:", err)
	}

	// Check the driver connections.
	if len(driverCons) != 2 {
		log.Fatalf("Expected 2 driver connections, but found %v.", len(driverCons))
	}

	srcDbDriverConn := driverCons[0]
	if srcDbDriverConn == nil {
		log.Fatal("The source database driver connection is nil.")
	}
	destDbDriverConn := driverCons[1]
	if destDbDriverConn == nil {
		log.Fatal("The destination database driver connection is nil.")
	}

	// Prepare to perform the backup.
	backup, err := destDbDriverConn.Backup("main", srcDbDriverConn, "main")
	if err != nil {
		log.Println(err)
	}

	isDone, err := backup.Step(0)
	if err != nil {
		log.Println(err)
	}
	if isDone {
		log.Println("Backup is unexpectedly done.")
	}

	// Check that the page count and remaining values are reasonable.
	initialPageCount := backup.PageCount()
	if initialPageCount <= 0 {
		log.Fatalf("Unexpected initial page count value: %v", initialPageCount)
	}
	initialRemaining := backup.Remaining()
	if initialRemaining <= 0 {
		log.Fatalf("Unexpected initial remaining value: %v", initialRemaining)
	}
	if initialRemaining != initialPageCount {
		log.Fatalf("Initial remaining value differs from the initial page count value; remaining: %v; page count: %v", initialRemaining, initialPageCount)
	}

	const usePagePerStepsTimeoutSeconds = 30
	usePerPageSteps := false

	// Perform the backup.
	if usePerPageSteps {
		var startTime = time.Now().Unix()

		// Test backing-up using a page-by-page approach.
		var latestRemaining = initialRemaining
		for {
			// Perform the backup step.
			isDone, err = backup.Step(1)
			if err != nil {
				log.Fatal("Failed to perform a backup step:", err)
			}

			// The page count should remain unchanged from its initial value.
			currentPageCount := backup.PageCount()
			if currentPageCount != initialPageCount {
				log.Fatalf("Current page count differs from the initial page count; initial page count: %v; current page count: %v", initialPageCount, currentPageCount)
			}

			// There should now be one less page remaining.
			currentRemaining := backup.Remaining()
			expectedRemaining := latestRemaining - 1
			if currentRemaining != expectedRemaining {
				log.Fatalf("Unexpected remaining value; expected remaining value: %v; actual remaining value: %v", expectedRemaining, currentRemaining)
			}
			latestRemaining = currentRemaining

			if isDone {
				break
			}

			// Limit the runtime of the backup attempt.
			if (time.Now().Unix() - startTime) > usePagePerStepsTimeoutSeconds {
				log.Fatal("Backup is taking longer than expected.")
			}
		}
	} else {
		// Test the copying of all remaining pages.
		isDone, err = backup.Step(-1)
		if err != nil {
			log.Fatal("Failed to perform a backup step:", err)
		}
		if !isDone {
			log.Fatal("Backup is unexpectedly not done.")
		}
	}

	// Check that the page count and remaining values are reasonable.
	if finalPageCount := backup.PageCount(); finalPageCount != initialPageCount {
		log.Fatalf("Final page count differs from the initial page count; initial page count: %v; final page count: %v", initialPageCount, finalPageCount)
	}
	if finalRemaining := backup.Remaining(); finalRemaining != 0 {
		log.Fatalf("Unexpected remaining value: %v", finalRemaining)
	}

	// Finish the backup.
	if err = backup.Finish(); err != nil {
		log.Fatal("Failed to finish backup:", err)
	}

	return srcDb, destDb
}

func TestDF(t *testing.T) {

	var driverCons []*sqlite3.SQLiteConn
	var driverName = "sqlite3_with_hook"

	sql.Register(driverName, &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			driverCons = append(driverCons, conn)
			return nil
		}})

	// Connect to the source database.
	srcTempFilename := "file:ent.db?mode=rwc&cache=shared&_fk=1&_journal=WAL&_cache_size=2000"

	srcDb, err := sql.Open(driverName, srcTempFilename)
	if err != nil {
		t.Fatal("Failed to open the source database:", err)
	}
	defer srcDb.Close()
	err = srcDb.Ping()
	if err != nil {
		t.Fatal("Failed to connect to the source database:", err)
	}

	// Connect to the destination database.
	destTempFilename := "file:aff.db?mode=memory&cache=shared&_fk=1&_cache_size=2000"

	destDb, err := sql.Open(driverName, destTempFilename)
	if err != nil {
		t.Fatal("Failed to open the destination database:", err)
	}
	defer destDb.Close()
	err = destDb.Ping()
	if err != nil {
		t.Fatal("Failed to connect to the destination database:", err)
	}

	// Check the driver connections.
	if len(driverCons) != 2 {
		t.Fatalf("Expected 2 driver connections, but found %v.", len(driverCons))
	}

	srcDbDriverConn := driverCons[0]
	if srcDbDriverConn == nil {
		t.Fatal("The source database driver connection is nil.")
	}
	destDbDriverConn := driverCons[1]
	if destDbDriverConn == nil {
		t.Fatal("The destination database driver connection is nil.")
	}

	var generateTestData = func(id int) string {
		return fmt.Sprintf("test-%v", id)
	}
	testRowCount := 100

	tx, err := srcDb.Begin()
	if err != nil {
		t.Fatal("Failed to begin a transaction when populating the source database:", err)
	}
	_, err = srcDb.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, value TEXT)")
	if err != nil {
		tx.Rollback()
		t.Fatal("Failed to create the source database \"test\" table:", err)
	}
	for id := 0; id < testRowCount; id++ {
		_, err = srcDb.Exec("INSERT INTO test (id, value) VALUES (?, ?)", id, generateTestData(id))
		if err != nil {
			tx.Rollback()
			t.Fatal("Failed to insert a row into the source database \"test\" table:", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal("Failed to populate the source database:", err)
	}

	// Prepare to perform the backup.
	backup, err := destDbDriverConn.Backup("main", srcDbDriverConn, "main")
	if err != nil {
		t.Fatal("Failed to initialize the backup:", err)
	}

	// Allow the initial page count and remaining values to be retrieved.
	// According to <https://www.sqlite.org/c3ref/backup_finish.html>, the page count and remaining values are "... only updated by sqlite3_backup_step()."
	isDone, err := backup.Step(0)
	if err != nil {
		t.Fatal("Unable to perform an initial 0-page backup step:", err)
	}
	if isDone {
		t.Fatal("Backup is unexpectedly done.")
	}

	// Check that the page count and remaining values are reasonable.
	initialPageCount := backup.PageCount()
	if initialPageCount <= 0 {
		t.Fatalf("Unexpected initial page count value: %v", initialPageCount)
	}
	initialRemaining := backup.Remaining()
	if initialRemaining <= 0 {
		t.Fatalf("Unexpected initial remaining value: %v", initialRemaining)
	}
	if initialRemaining != initialPageCount {
		t.Fatalf("Initial remaining value differs from the initial page count value; remaining: %v; page count: %v", initialRemaining, initialPageCount)
	}

	const usePagePerStepsTimeoutSeconds = 30
	usePerPageSteps := false

	// Perform the backup.
	if usePerPageSteps {
		var startTime = time.Now().Unix()

		// Test backing-up using a page-by-page approach.
		var latestRemaining = initialRemaining
		for {
			// Perform the backup step.
			isDone, err = backup.Step(1)
			if err != nil {
				t.Fatal("Failed to perform a backup step:", err)
			}

			// The page count should remain unchanged from its initial value.
			currentPageCount := backup.PageCount()
			if currentPageCount != initialPageCount {
				t.Fatalf("Current page count differs from the initial page count; initial page count: %v; current page count: %v", initialPageCount, currentPageCount)
			}

			// There should now be one less page remaining.
			currentRemaining := backup.Remaining()
			expectedRemaining := latestRemaining - 1
			if currentRemaining != expectedRemaining {
				t.Fatalf("Unexpected remaining value; expected remaining value: %v; actual remaining value: %v", expectedRemaining, currentRemaining)
			}
			latestRemaining = currentRemaining

			if isDone {
				break
			}

			// Limit the runtime of the backup attempt.
			if (time.Now().Unix() - startTime) > usePagePerStepsTimeoutSeconds {
				t.Fatal("Backup is taking longer than expected.")
			}
		}
	} else {
		// Test the copying of all remaining pages.
		isDone, err = backup.Step(-1)
		if err != nil {
			t.Fatal("Failed to perform a backup step:", err)
		}
		if !isDone {
			t.Fatal("Backup is unexpectedly not done.")
		}
	}

	// Check that the page count and remaining values are reasonable.
	finalPageCount := backup.PageCount()
	if finalPageCount != initialPageCount {
		t.Fatalf("Final page count differs from the initial page count; initial page count: %v; final page count: %v", initialPageCount, finalPageCount)
	}
	finalRemaining := backup.Remaining()
	if finalRemaining != 0 {
		t.Fatalf("Unexpected remaining value: %v", finalRemaining)
	}

	// Finish the backup.
	err = backup.Finish()
	if err != nil {
		t.Fatal("Failed to finish backup:", err)
	}

	desRowCount := 200

	txd, err := destDb.Begin()
	if err != nil {
		t.Fatal("Failed to begin a transaction when populating the source database:", err)
	}
	_, err = destDb.Exec("CREATE TABLE test2 (id INTEGER PRIMARY KEY, value TEXT)")
	if err != nil {
		txd.Rollback()
		t.Fatal("Failed to create the source database \"test\" table:", err)
	}
	for id := 100; id < desRowCount; id++ {
		_, err = destDb.Exec("INSERT INTO test2 (id, value) VALUES (?, ?)", id, generateTestData(id))
		if err != nil {
			txd.Rollback()
			t.Fatal("Failed to insert a row into the source database \"test\" table:", err)
		}
	}
	err = txd.Commit()
	if err != nil {
		t.Fatal("Failed to populate the source database:", err)
	}

	rows, err := destDb.Query("select id,value,res from ( select id,value, 1 as res from test union all select id,value,-1 as res from test2) group by id,value having sum(res);")

	var res []Ttbale
	for rows.Next() {
		var s Ttbale
		if err := rows.Scan(&s.Id, &s.Value, &s.Has); err != nil {
			log.Fatal(err)
		}
		res = append(res, s)
	}
	log.Println(res)

}

func TestRead(t *testing.T) {
	const usePagePerStepsTimeoutSeconds = 30
	driverConns := []*sqlite3.SQLiteConn{}
	var driverName = "sqlite3_with_hook_example"

	sql.Register(driverName, &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			driverConns = append(driverConns, conn)
			return nil
		},
	})

	// Connect to the source database.
	srcTempFilename := "file:ent.db?mode=rwc&cache=shared&_fk=1&_journal=WAL&_cache_size=2000"

	srcDb, err := sql.Open(driverName, srcTempFilename)
	if err != nil {
		t.Fatal("Failed to open the source database:", err)
	}
	defer srcDb.Close()
	err = srcDb.Ping()
	if err != nil {
		t.Fatal("Failed to connect to the source database:", err)
	}

	// Connect to the destination database.
	destTempFilename := "file:aff.db?mode=memory&cache=shared&_fk=1&_cache_size=2000"

	destDb, err := sql.Open(driverName, destTempFilename)
	if err != nil {
		t.Fatal("Failed to open the destination database:", err)
	}
	defer destDb.Close()
	err = destDb.Ping()
	if err != nil {
		t.Fatal("Failed to connect to the destination database:", err)
	}

	// Check the driver connections.
	if len(driverConns) != 2 {
		t.Fatalf("Expected 2 driver connections, but found %v.", len(driverConns))
	}

	srcDbDriverConn := driverConns[0]
	if srcDbDriverConn == nil {
		t.Fatal("The source database driver connection is nil.")
	}
	destDbDriverConn := driverConns[1]
	if destDbDriverConn == nil {
		t.Fatal("The destination database driver connection is nil.")
	}

	// Prepare to perform the backup.
	backup, err := destDbDriverConn.Backup("main", srcDbDriverConn, "main")
	if err != nil {
		t.Fatal("Failed to initialize the backup:", err)
	}

	// Allow the initial page count and remaining values to be retrieved.
	// According to <https://www.sqlite.org/c3ref/backup_finish.html>, the page count and remaining values are "... only updated by sqlite3_backup_step()."
	isDone, err := backup.Step(0)
	if err != nil {
		t.Fatal("Unable to perform an initial 0-page backup step:", err)
	}
	if isDone {
		t.Fatal("Backup is unexpectedly done.")
	}

	// Check that the page count and remaining values are reasonable.
	initialPageCount := backup.PageCount()
	if initialPageCount <= 0 {
		t.Fatalf("Unexpected initial page count value: %v", initialPageCount)
	}
	initialRemaining := backup.Remaining()
	if initialRemaining <= 0 {
		t.Fatalf("Unexpected initial remaining value: %v", initialRemaining)
	}
	if initialRemaining != initialPageCount {
		t.Fatalf("Initial remaining value differs from the initial page count value; remaining: %v; page count: %v", initialRemaining, initialPageCount)
	}

	//usePerPageSteps := false
	//
	//// Perform the backup.
	//if usePerPageSteps {
	//	var startTime = time.Now().Unix()
	//
	//	// Test backing-up using a page-by-page approach.
	//	var latestRemaining = initialRemaining
	//	for {
	//		// Perform the backup step.
	//		isDone, err = backup.Step(1)
	//		if err != nil {
	//			t.Fatal("Failed to perform a backup step:", err)
	//		}
	//
	//		// The page count should remain unchanged from its initial value.
	//		currentPageCount := backup.PageCount()
	//		if currentPageCount != initialPageCount {
	//			t.Fatalf("Current page count differs from the initial page count; initial page count: %v; current page count: %v", initialPageCount, currentPageCount)
	//		}
	//
	//		// There should now be one less page remaining.
	//		currentRemaining := backup.Remaining()
	//		expectedRemaining := latestRemaining - 1
	//		if currentRemaining != expectedRemaining {
	//			t.Fatalf("Unexpected remaining value; expected remaining value: %v; actual remaining value: %v", expectedRemaining, currentRemaining)
	//		}
	//		latestRemaining = currentRemaining
	//
	//		if isDone {
	//			break
	//		}
	//
	//		// Limit the runtime of the backup attempt.
	//		if (time.Now().Unix() - startTime) > usePagePerStepsTimeoutSeconds {
	//			t.Fatal("Backup is taking longer than expected.")
	//		}
	//	}
	//} else {
	//	// Test the copying of all remaining pages.
	//	isDone, err = backup.Step(-1)
	//	if err != nil {
	//		t.Fatal("Failed to perform a backup step:", err)
	//	}
	//	if !isDone {
	//		t.Fatal("Backup is unexpectedly not done.")
	//	}
	//}

	// Check that the page count and remaining values are reasonable.
	finalPageCount := backup.PageCount()
	if finalPageCount != initialPageCount {
		t.Fatalf("Final page count differs from the initial page count; initial page count: %v; final page count: %v", initialPageCount, finalPageCount)
	}
	finalRemaining := backup.Remaining()
	if finalRemaining != 0 {
		t.Fatalf("Unexpected remaining value: %v", finalRemaining)
	}

	// Finish the backup.
	err = backup.Finish()
	if err != nil {
		t.Fatal("Failed to finish backup:", err)
	}

	//rows, err := destDb.Query("select * from test")
	//
	//var res []Ttbale
	//for rows.Next() {
	//	var s Ttbale
	//	if err := rows.Scan(&s.Id, &s.Value); err != nil {
	//		log.Fatal(err)
	//	}
	//	res = append(res, s)
	//}
	//log.Println(res)
}

func TestMemDB(t *testing.T) {

	const usePagePerStepsTimeoutSeconds = 30
	driverConns := []*sqlite3.SQLiteConn{}
	var driverName = "sqlite3_with_hook_example"

	sql.Register(driverName, &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			driverConns = append(driverConns, conn)
			return nil
		},
	})

	destTempFilename := "file:aff.db?mode=memory&cache=shared&_fk=1&_cache_size=2000"

	srcDb, err := sql.Open(driverName, destTempFilename)
	if err != nil {
		t.Fatal("Failed to open the destination database:", err)
	}

	testRowCount := 1000000

	var generateTestData = func(id int) string {
		return fmt.Sprintf("test-%v", id)
	}
	tx, err := srcDb.Begin()
	if err != nil {
		t.Fatal("Failed to begin a transaction when populating the source database:", err)
	}
	_, err = srcDb.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, value TEXT)")
	if err != nil {
		tx.Rollback()
		t.Fatal("Failed to create the source database \"test\" table:", err)
	}
	for id := 0; id < testRowCount; id++ {
		_, err = srcDb.Exec("INSERT INTO test (id, value) VALUES (?, ?)", id, generateTestData(id))
		if err != nil {
			tx.Rollback()
			t.Fatal("Failed to insert a row into the source database \"test\" table:", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal("Failed to populate the source database:", err)
	}
}
