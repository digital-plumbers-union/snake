package buildnum

// BuildNumber is a struct that represents the current state of our build number, so that
// the webhook server can update it independently of the controller which syncs it with a configmap
type BuildNumber struct {
	Number int64
}

// New instantiates BuildNumber objects
func New(number int64) BuildNumber {
	return BuildNumber{
		Number: number,
	}
}

// Increment increases the build number by 1
func (b *BuildNumber) Increment() {
	b.Number = b.Number + 1
}

// Get retrieves the current build number
func (b BuildNumber) Get() int64 {
	return b.Number
}

// Set allows setting current build number
func (b *BuildNumber) Set(number int64) {
	b.Number = number
}
