module github.com/d-kuro/restart-object

go 1.12

require (
	github.com/kubernetes/kubernetes v1.14.3
	github.com/spf13/cobra v0.0.5
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	k8s.io/api v0.0.0-20190602205700-9b8cae951d65
	k8s.io/apimachinery v0.0.0-20190602125621-c0632ccbde11
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/utils v0.0.0-20190607212802-c55fbcfc754a // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190602130007-e65ca70987a6
