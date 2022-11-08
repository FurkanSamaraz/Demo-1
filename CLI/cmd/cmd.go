package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd, uygulama alt-komutsuz şekilde çağırıldığında
// gerçekleşecekleri temsil eder.
var rootCmd = &cobra.Command{
	Use:   "a",
	Short: "uygulamanın kısa bilgisi",
	Long:  `uygulamanın uzun bilgilendirme metnini bu kısma girebilirsinizaaa.`,
}

// Execute fonksiyonu tüm komut ve alt-komutları
// uygulama çalıştırıldığında yüklenmesini sağlar
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// init fonksiyonu içerisinde ise komuta bağlı olan
	// alt-komutlarını ve flag'lerini bağlayabilir ve bunlar
	// için özelleştirmelerde bulunabilirsiniz.
	rootCmd.Flags().BoolP("toggle", "t", false, "toggle için yardım mesajı")
}
