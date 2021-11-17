package main

import (
	"fmt"
	"os"

	"go.m3o.com/user"
)

// Send a verification email
// to the user being signed up. Email from will be from 'noreply@email.m3ocontent.com',
// but you can provide the title and contents.
// The verification link will be injected in to the email as a template variable, $micro_verification_link.
// Example: 'Hi there, welcome onboard! Use the link below to verify your email: $micro_verification_link'
// The variable will be replaced with an actual url that will look similar to this:
// 'https://user.m3o.com/user/verify?token=a-verification-token&redirectUrl=your-redir-url'
func main() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.SendVerificationEmail(&user.SendVerificationEmailRequest{
		Email:              "joe@example.com",
		FailureRedirectUrl: "https://m3o.com/verification-failed",
		FromName:           "Awesome Dot Com",
		RedirectUrl:        "https://m3o.com",
		Subject:            "Email verification",
		TextContent: `Hi there,

Please verify your email by clicking this link: $micro_verification_link`,
	})
	fmt.Println(rsp, err)

}
