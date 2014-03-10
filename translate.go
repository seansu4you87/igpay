package igpay

const (
	pigLatinSuffix string = "ay"
)

func Translate(in string) (out string) {
	out = in[1:] + in[0:1] + pigLatinSuffix
	return
}
