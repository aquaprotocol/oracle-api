/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package deposition

// Klasa zawierająca dane przelewu / Transfer Data Class
type TransferDataCurrencyRequired struct {

	// Waluta, domyślnie EUR przelewów zagranicznych / Currency
	Currency string `json:"currency"`

	// Pole opisujące przelew / Description
	Description string `json:"description"`

	// Kwota przelewu / Amount
	Amount string `json:"amount"`

	// Data wykonania przelewu. Wymagany pod warunkiem przekazania wartości FutureDated w atrybucie executionMode. Format: YYYY-MM-DD / Date when the transfer is to be executed (YYYY-MM-DD). Required conditionally if executionMode attribute has value FutureDated.
	ExecutionDate string `json:"executionDate,omitempty"`
}
