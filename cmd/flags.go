package cmd

type Options struct {
	JwtToken string `short:"t" long:"token" description:"JWT token to crack" required:"true"`
	List     string `short:"l" long:"list" description:"Specify the file dict" required:"true"`
}
