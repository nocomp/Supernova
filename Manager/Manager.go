package Manager

import (
	"Supernova/Arguments"
	"Supernova/Converters"
	"Supernova/Encryptors"
	"Supernova/Output"
	"fmt"
	"strings"
)

// EncryptionManager function
func EncryptionManager(Key int, Encryption string, Obfuscation string, Variable string, rawShellcode string, foundLanguage string, fileSizeFlag bool) (string, string) {
	// Call function ValidateKeySize
	Key = Arguments.ValidateKeySize(Key, Encryption)

	// Call function named DetectEncryption
	encryptedShellcode, encryptedLength := Encryptors.DetectEncryption(Encryption, rawShellcode, Key, foundLanguage)

	// Call function named ConvertShellcode2Template
	template := Converters.ConvertShellcode2Template(encryptedShellcode, foundLanguage, encryptedLength, Variable)

	// if Obfuscation is empty
	if Obfuscation == "" {
		switch fileSizeFlag {
		// if fileSizeFlag is true
		case true:
			fmt.Printf("[!] The size of the encrypted shellcode exceeds the maximum display limit.\n\n[!] Supernova cannot display it on the screen.\n\n")
		default:
			if foundLanguage == "raw" {
				fmt.Printf("[!] The encrypted shellcode is displayed in raw format represented as hexadecimal on the terminal.\n\n")
			}
			// Print encrypted template
			fmt.Printf("[+] The encrypted payload with %s:\n\n%s\n\n", strings.ToUpper(Encryption), template)
		}
	}

	return template, encryptedShellcode
}

// OutputManager function
func OutputManager(OutFile string, Language string, template string) {
	// Outfile option is enable
	if OutFile != "" {
		language := strings.ToLower(Language)
		if language == "raw" {
			err := Output.SaveShellcodeToFile(template, OutFile)
			if err != nil {
				fmt.Println("[!] Error:", err)
				return
			}
		} else {
			err := Output.SaveOutputToFile(template, OutFile)
			if err != nil {
				fmt.Println("[!] Error:", err)
				return
			}
		}
	}
}
