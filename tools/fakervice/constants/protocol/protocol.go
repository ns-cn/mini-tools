package protocol

type Protocol int

const (
	TCP Protocol = iota
	UDP
	HTTP
	HTTPS
)
const (
	tcp = ""
)

func (p Protocol) String() string {
	switch p {
	case TCP:
		return "TCP"
	case UDP:
		return "UDP"
	case HTTP:
		return "HTTP"
	case HTTPS:
		return "HTTPS"
	}
	return ""
}
