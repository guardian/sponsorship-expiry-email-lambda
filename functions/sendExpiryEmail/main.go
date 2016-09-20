package main

import (
	//"github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"html/template"
	"os"
	"time"
	"fmt"
)

const emailTemplate = `
<!DOCTYPE html>
<html>
	<body>
		<h3>Expiring sponsorships</h3>
		{{range .Expiring}}
			<div>
			<span>{{with $asset := index .SponsorLogo.Assets 0}}<img src="{{$asset.ImageUrl}}"/>{{end}}</span>
			<span>
				{{.SponsorshipType}}<br />
				Sponsor: <a href="https://tagmanager.gutools.co.uk/sponsorship/{{.Id}}">{{.SponsorName}}</a><br />
				Expires: {{formatMillisDate .ValidTo}}
			</span>
			</div>
		{{else}}
			<p>no sponsorships expiring in the next 7 days</p>
		{{end}}

		<h3>Expired sponsorships</h3>
		{{range .Expired}}
			<div>
			<span>{{with $asset := index .SponsorLogo.Assets 0}}<img src="{{$asset.ImageUrl}}"/>{{end}}</span>
			<span>
				{{.SponsorshipType}}<br />
				Sponsor: <a href="https://tagmanager.gutools.co.uk/sponsorship/{{.Id}}">{{.SponsorName}}</a><br />
				Expired: {{formatMillisDate .ValidTo}}
			</span>
			</div>
		{{else}}
			<p>no sponsorships expired in the last 7 days</p>
		{{end}}
	</body>
</html>`

func formatMillisDate(millis int64) string {
	return time.Unix(millis / 1000, 0).Format(time.RFC1123)
}

func main() {
	dynamo := dynamodb.New(session.New(&aws.Config{
		Region: aws.String("eu-west-1"),
	}))

	expiring, err := LoadExpiringSoon(dynamo)

	if (err != nil) {
		return
	}

	expired, err := LoadExpiredRecently(dynamo)

	if (err != nil) {
		return
	}

	t, err := template.New("webpage").Funcs(template.FuncMap{"formatMillisDate": formatMillisDate,}).Parse(emailTemplate)

	templateData := struct {
		Expiring []Sponsorship
		Expired []Sponsorship
	}{
		Expiring: expiring,
		Expired: expired,
	}

	err = t.Execute(os.Stdout, templateData)

	if (err != nil) {
		fmt.Println(err.Error())
	}

	//apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
	//	dynamo := dynamodb.New(session.New())
	//
	//	LoadExpiringSoon(dynamo)
	//
	//	return m, nil
	//})
}