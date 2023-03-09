package numberinfull

import (
	"strings"
)

var (
	de1a9 = []string{
		"", "um", "dois", "tres", "quatro", "cinco", "seis", "sete", "oito",
		"nove",
	}
	de10a19 = []string{
		"dez", "onze", "doze", "treze", "quatorze", "quinze", "dezeseis",
		"dezesete", "dezoito", "dezenove",
	}
	de20a90 = []string{
		"", "", "vinte ", "trinta ", "quarenta ", "cinquenta ", "sessenta ",
		"setenta ", "oitenta ", "noventa ",
	}
	de100a900 = []string{
		"", "cento ", "duzentos ", "trezentos ", "quatrocentos ", "quinhentos ",
		"seiscentos ", "setecentos ", "oitocentos ", "novecentos ",
	}
)

func numberINFull(n int32) string {
	var output string

	if n >= 100000 {
		return "numero fora do limite"
	}

	if n >= 20000 && n < 100000 {
		output += de20a90[n/10000]
		if (n/1000)%10 == 0 {
			output += "mil "
		}
		n %= 10000
	}

	if n >= 10000 && n < 20000 {
		output += de10a19[n/1000-10] + " mil "
		n %= 1000
	}

	if n >= 1000 && n < 10000 {
		if len(output) > 0 {
			output += "e "
		}
		if n >= 2000 {
			output += de1a9[n/1000] + " "
		}
		output += "mil "
		n %= 1000
	}

	if n >= 100 && n < 1000 {
		if n == 100 {
			if len(output) > 0 {
				output += "e "
			}
			output += "cem"
		} else {
			if len(output) > 0 && n%100 == 0 {
				output += "e "
			}
			output += de100a900[n/100]
		}
		n %= 100
	}

	if n >= 20 && n < 100 {
		if len(output) > 0 {
			output += "e "
		}
		output += de20a90[n/10]
		n %= 10
	}

	if n >= 10 && n < 20 {
		if len(output) > 0 {
			output += "e "
		}
		output += de10a19[n-10]
		n -= 20
	}

	if n > 0 && n < 10 {
		if len(output) > 0 {
			output += "e "
		}
		output += de1a9[n]
	}

	return strings.TrimSpace(output)
}
