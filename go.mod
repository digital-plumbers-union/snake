module github.com/digital-plumbers-union/snake

go 1.15

require (
	github.com/tektoncd/pipeline v0.15.2
	golang.org/x/tools v0.0.0-20200616195046-dc31b401abb5
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	sigs.k8s.io/controller-runtime v0.5.10
	sigs.k8s.io/controller-tools v0.4.0
	sigs.k8s.io/kind v0.8.1
)

// pin to 0.17.6 for compatibility with Tekton
replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.2.9
)
