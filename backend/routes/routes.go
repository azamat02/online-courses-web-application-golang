package routes

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/controllers"
)

func Setup(app *fiber.App) {
	//User routes
	app.Get("api/users/:id", controllers.GetUser)
	app.Get("api/users", controllers.GetAllUsers)

	//Authorization routes
	app.Post("api/sign_up", controllers.SignUp)
	app.Post("api/sign_in", controllers.SignIn)
	app.Get("api/user", controllers.User)
	app.Get("api/admin", controllers.IsAdmin)
	app.Post("api/logout", controllers.Logout)

	//Courses routes
	app.Get("api/courses/:id", controllers.GetCourse)
	app.Get("api/courses", controllers.GetAllCourses)
	app.Post("api/courses", controllers.CreateCourse)

	//Module routes
	app.Get("api/modules/:id", controllers.GetModule)
	app.Get("api/modules/by_course/:id", controllers.GetModulesByCourseId)
	app.Get("api/modules", controllers.GetAllModules)
	app.Post("api/modules", controllers.CreateModule)

	//Lesson routes
	app.Get("api/lessons/:id", controllers.GetLesson)
	app.Get("api/lessons/by_module/:id", controllers.GetLessonsByModuleId)
	app.Get("api/lessons", controllers.GetAllLessons)
	app.Post("api/lessons", controllers.CreateLesson)

	//Comment routes
	app.Get("api/comments/:id", controllers.GetComment)
	app.Get("api/comments/by_course/:id", controllers.GetCommentsByCourseId)
	app.Get("api/comments", controllers.GetAllComments)
	app.Post("api/comments", controllers.CreateComment)

	//Purchased courses routes
	app.Get("api/purchased/", controllers.GetAllPurchasedCourses)
	app.Get("api/purchased/by_user/:id", controllers.GetAllUserPurchasedCourses)
	app.Get("api/purchased/has/:id", controllers.CheckIfUserHasPurchasedCourse)
	app.Post("api/purchased", controllers.CreatePurchasedCourse)

	//Admin routes
	app.Post("api/admin/create_user", controllers.SignUp)
	app.Post("api/admin/delete_user", controllers.DeleteUserById)
	app.Post("api/admin/update_user", controllers.UpdateUserById)

	app.Post("api/admin/create_course", controllers.CreateCourse)
	app.Post("api/admin/delete_course", controllers.DeleteCourseById)
	app.Post("api/admin/update_course", controllers.UpdateCourseById)

	app.Post("api/admin/create_module", controllers.CreateModule)
	app.Post("api/admin/delete_module", controllers.DeleteModuleById)
	app.Post("api/admin/update_module", controllers.UpdateModuleById)

	app.Post("api/admin/create_lesson", controllers.CreateLesson)
	app.Post("api/admin/delete_lesson", controllers.DeleteLessonById)
	app.Post("api/admin/update_lesson", controllers.UpdateLessonById)
}
