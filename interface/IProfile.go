//5 method signatures.
/*createfile
 editfile
 delete profile
 fetch profile*/

 //this interface will proviode the required methods
 package interfaces

 import "rest-api/models"

 type IProfile interface{
	CreateProfile(profile *models.Profile)
	EditProfile(profile *models.Profile)
	DeleteProfile(profileId int)
	GetProfileById(profile int)
	GetProfileBySearch()
 }