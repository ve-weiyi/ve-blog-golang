package https

import (
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

func TestNewHttpBuilder(t *testing.T) {
	get, err := NewHttpBuilder("https://test_saas.qdtech.ai/index.php/api/survey/getSurveyList?access_key=K%2FoKT09jjFHDj%2B6xXVydgCOBpemSFMGQ9NEfPXgci%2FfdNLhXwZ3by6UO%2B1%2BtfREB%2Bz5LJ%2FuzQOn1rL6J%2FPk3A7hEfvkAbvAsLpXChZYmHTUTr4N325ouinb24gVAfKA3ybJ%2FWZZs1MTYz1AqZP51C7tO2zLDcy%2BxtbO1Ifz84wu2LIwxKeB%2BzY16AEuRHgXQPoodj7petWHVuOpmvj8BTVBMXVtE3rb3zLdX72ha9q1SXEU3oHl9VwCbmw0uorYT7Qj06XLkN%2BmiTZIbYTETqzi0hEVlf0j5kjUg3smuWDV7D7VdHQ0vrtaz7hwwbh8rcIrYikvTO2ozM2QvNB%2FMWdEsnyeEZTIWlH1uWSlui5TUzvddZchsE8FKSzCCrDU4").
		Get()
	if err != nil {
		t.Log(err)
	}
	t.Log(jsonconv.ObjectToJsonIndent(get))
}
