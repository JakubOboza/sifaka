package display

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/JakubOboza/sifaka/storage"
)

func PrintCertificateDataList(certDatums *[]storage.CertificateData) {
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	header := []string{
		"id", "name", "subject", "type", "not_after", "expires_in", "status",
	}

	if err := w.Write(header); err != nil {
		fmt.Println(err)
		return
	}

	for _, certData := range *certDatums {
		record := []string{
			fmt.Sprintf("%d", certData.ID),
			certData.Name,
			certData.Subject,
			certData.Type,
			certData.NotAfter.String(),
			certData.ExpiresInString(),
			certData.Status(),
		}
		if err := w.Write(record); err != nil {
			fmt.Println(err)
			return
		}
	}

}
