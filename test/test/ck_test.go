package test

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DirItemMeta struct {
	Id      string
	Fol     string
	Pfol    string
	Step    uint
	Isdir   uint
	Md5     string
	Modtime int64
	Res     int
}

func CompareLocalSyncPathChange(sqlConn *sqlx.DB, tab1, tab2 string) {

	sprintf := fmt.Sprintf(
		"select id,fol,pfol,step,isdir,md5,modtime,res from ( select id,fol,pfol,step,isdir,md5,modtime, 1 as res from %s union all select id,fol,pfol,step,isdir,md5,modtime,-1 as res from %s) group by fol,pfol,step,modtime having sum(res);", tab1, tab2)
	rows, err := sqlConn.Queryx(sprintf)
	if err != nil {
		log.Println(err)
	}

	//diff
	//rows, err := sqlConn.Queryx("select fol,pfol,step,times,res from ( select fol,pfol,step,times, 1 as res from oldfile union all select fol,pfol,step,times,-1 as res from newfile) group by fol,pfol,step,times having sum(res);")
	//if err != nil {
	//	log.Printf("%v", err)
	//}

	log.Println("-----------------")

	//var cs []DirItemMeta
	for rows.Next() {
		var f DirItemMeta
		rows.StructScan(&f)

		sta := "add"
		if f.Res > 0 {
			sta = "update"
		}

		log.Println(f, sta)

		//cs = append(cs, f)
		//file[f.Id] = &FileChange{
		//	ID:           f.Id,
		//	State:        sta,
		//	Folder:       f.Fol,
		//	ParentFolder: f.Pfol,
		//	Step:         f.Step,
		//	Times:        f.Times,
		//}
		//log.Println(f)
	}

	//log.Println("-----------------")
	//
	//rows, err = mem.Queryx("select id,fol,pfol,step,res from ( select id,fol,pfol,step, 1 as res from oldfile union all select id,fol,pfol,step,-1 as res from newfile) group by fol,pfol,step having sum(res);")
	//if err != nil {
	//	log.Printf("%v", err)
	//}
	//
	//for rows.Next() {
	//	var f Mem
	//	rows.StructScan(&f)
	//	if f.Res > 0 {
	//		file[f.Id].State = "delete"
	//	}
	//	log.Println(f)
	//}
	//
	//for id, change := range file {
	//	log.Println(id, change.Folder, change.State)
	//}
	//
	//sprintf := fmt.Sprintf(
	//	"select id,fol,pfol,step,isdir,md5,res from ( select id,fol,pfol,step,isdir,md5, 1 as res from %s union all select id,fol,pfol,step,isdir,md5,-1 as res from %s) group by fol,pfol,step having sum(res);", tab1, tab2)
	//rows, err := sqlConn.Query(sprintf)
	//if err != nil {
	//	log.Println("232", err)
	//}
	//
	//var res []DirItemMeta
	//for rows.Next() {
	//	var s DirItemMeta
	//	if err := rows.Scan(&s.Id, &s.Fol, &s.Pfol, &s.Step, &s.Isdir, &s.Md5, &s.Res); err != nil {
	//		log.Fatal(err)
	//	}
	//	res = append(res, s)
	//}
	//log.Println(res)

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
	//
	//CompareLocalSyncPathChange(destDb, "ts1", "ts2")

}

// 遍历路径至数据库
func WalkDirToDataBase(sqlConn *sqlx.DB, Path, tab string) error {
	createTableSql := fmt.Sprintf("CREATE TABLE %s (id TEXT,fol TEXT,pfol TEXT,step INTEGER,isdir INTEGER,md5 TEXT,modtime INTEGER)", tab)
	insertTableSql := fmt.Sprintf("INSERT INTO %s (id,fol, pfol,step,isdir,md5) VALUES (?,?,?,?,?,?,?)", tab)

	p, _ := os.Stat(Path)

	if _, err := sqlConn.Exec(createTableSql); err != nil {
		log.Fatal(err)
	}

	dir, file := filepath.Split(Path)
	if _, err := sqlConn.Exec(insertTableSql, "", file, dir, 0, 1, "", p.ModTime()); err != nil {
		log.Fatal(err)
	}
	dirlen := len(dir)
	pathSeparator := string(os.PathSeparator)

	err := WalkDir(Path, func(path string, d fs.FileInfo, err error) error {
		split := strings.Split(path[dirlen:], pathSeparator)
		splen := len(split)

		if d.IsDir() {
			_, err = sqlConn.Exec(insertTableSql, "", split[splen-1], split[splen-2], splen-1, 1, "", d.ModTime())
		} else {
			_, err = sqlConn.Exec(insertTableSql, "", split[splen-1], split[splen-2], splen-1, 0, "", d.ModTime())
		}

		//time.Sleep(time.Second)
		//log.Printf(path, split[splen-2], split[splen-1], splen-1)

		return err
	})

	return err
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
