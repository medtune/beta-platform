package data

type Error struct {
	Main
	Code    int
	Message string
	Errors  string
	Details string
}

var (
	Error404 = &Error{
		Main: Main{
			PageTitle: "Error 404",
		},
		Code:    404,
		Message: "Oooops !",
		Details: "Page not found.",
	}

	Error405 = Error{
		Main: Main{
			PageTitle: "Error 405",
		},
		Code:    405,
		Message: "Method not allowed",
		Details: "Method not allowed",
	}

	Error401 = &Error{
		Main: Main{
			PageTitle: "Error 401",
		},
		Code:    401,
		Message: "Not too fast !",
		Details: "Unauthorized",
	}

	Error500 = Error{
		Main: Main{
			PageTitle: "Error 500",
		},
		Code:    500,
		Message: "Oooops !",
		Details: "Server error.",
	}

	Error501 = Error{
		Main: Main{
			PageTitle: "Error 500",
		},
		Code:    501,
		Message: "Not implemented",
		Details: "Not implemented",
	}

	ErrorImAteaPot = Error{
		Main: Main{
			PageTitle: "Error 418",
		},
		Code:    418,
		Message: "Wuuuuttt",
		Details: "Im a tea pot !",
	}

	ErrorFinalBoss = Error{
		Main: Main{
			PageTitle: "Error 418",
		},
		Code:    418,
		Message: "Final boss error",
		Details: "Dont try this at home!",
	}
)
