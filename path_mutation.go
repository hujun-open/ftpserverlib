package ftpserver

import "strings"

//ParamMutationFunc is a function modify in and return the modified string,
//cmd is the corresponding ftp command
type ParamMutationFunc func(cmd, inparam string) (outparam string)

//MakePathHandler returns a ParamMutationFunc that mutate parameter using mut for all commands has pathname as parameter in rfc959
func MakePathHandler(mut func(pathname string) string) ParamMutationFunc {
	return func(cmd, in string) string {
		switch cmd {
		case "CWD", "SMNT", "RETR", "STOR", "APPE", "RNFR", "RNTO", "DELE", "RMD", "MKD", "LIST", "NLST", "STAT":
			if strings.TrimSpace(in) != "" {
				return mut(in)
			}
		}
		return in
	}
}
