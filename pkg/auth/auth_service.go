package auth

//type Service struct {
//	jwt  jwt.Service
//	salt salt.Service
//	ur   domain.UserRepository
//}
//
//func (a *Service) Login(ctx context.Context, email, reqPassword string) (*ent.User, error) {
//	user, err := a.ur.GetByEmail(ctx, email)
//	comp := a.ComparePassword(reqPassword, user.Salt, user.PassWord)
//
//}
//
//func (a *Service) ComparePassword(reqPassword string, salt []byte, hashedPassword string) bool {
//	hashed := a.salt.Hashed([]byte(reqPassword), salt)
//	return hashed == hashedPassword
//}
