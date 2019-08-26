/**
 * Author：Robert WYQ
 * Date：2019-7-22 ——  2019-8-31
 * Description: Constants used in the TP | name should be consistent with the Client cli
 * ------------------------------------------------------------------------------
 */
package constants

const (
	// Related to transaction family
	TransactionFamilyName    string = "onlinestore"
	TransactionFamilyVersion string = "2.0"

	// Related to length of addresses, total is 70
	TransactionFamilyNamespaceAddressLength int = 6
	TransactionUserAddressLength            int = 64

	// Constants used in CLI
	TransactionTransport string = "transport"
	TransactionBuy  string = "buy"
	TransactionSell string = "sell"
	TransactionEmpty string = "empty"

	// Transaction indices
	TransactionOperationIndex int = 0
	TransactionAmountIndex    int = 1
	TransactionToIndex        int = 2
)
