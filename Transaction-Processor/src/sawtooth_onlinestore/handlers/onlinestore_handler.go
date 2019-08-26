/**
 * Author：Robert WYQ
 * Date：2019-7-22 ——  2019-8-31
 * Description: To create an onlinestore APP base on the sawtooth1.1.5
 * ------------------------------------------------------------------------------
 */

package handlers

import (
	"errors"
	"github.com/hyperledger/sawtooth-sdk-go/logging"
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/processor_pb2"
	"../constants"
	"../utils"
	"strconv"
	"strings"
)

// declare a log
var logger = logging.Get()

// get transaction family name
func (self *OnlinestoreHandler) FamilyName() string {
	return constants.TransactionFamilyName
}
// get transaction family version
func (self *OnlinestoreHandler) FamilyVersions() []string {
	return []string{constants.TransactionFamilyVersion}
}
// get hex code in string
func (self *OnlinestoreHandler) Namespaces() []string {
	return []string{self.getNamespaceAddress()}
}
// translate family name into hash hex code（prior 6）
func (self OnlinestoreHandler) getNamespaceAddress() string {
	return utils.Hex_encryption(constants.TransactionFamilyName)[:constants.TransactionFamilyNamespaceAddressLength]
}

// apply gets called with two argument：context 和 request， return with error
func (self *OnlinestoreHandler) Apply(request *processor_pb2.TpProcessRequest, context *processor.Context) error {
	// get payload from request， which send by validator
	payload := string(request.GetPayload())
	// parse the content of payload，divide into string[]
	payloadList := strings.Split(payload, ",")
	// examine the input payload action parameters
	if len(payloadList) != 2 && !(len(payloadList) == 3 && constants.TransactionTransport == payloadList[constants.TransactionOperationIndex]) {
		return errors.New("Invalid num of arguments: expected 2 or 3, got: " + string(len(payloadList)))
	}

	// put the input parameters into onlinestore handler
	self.context = context
	// get operation
	self.operation = payloadList[constants.TransactionOperationIndex]
	// get amount
	var err error
	self.amount, err = strconv.Atoi(payloadList[constants.TransactionAmountIndex])
	if err != nil {
		return err
	}
	// get public key from header
	self.AddressFrom = request.GetHeader().GetSignerPublicKey()

	// according to different operations, choose different actions
	switch self.operation {
	case constants.TransactionBuy:
		return self.buy()
	case constants.TransactionSell:
		return self.sell()
	case constants.TransactionTransport:
		self.AddressTo = payloadList[constants.TransactionToIndex]
		return self.transport()
	case constants.TransactionEmpty:
		return self.empty()
	default:
		return errors.New("Unsupported operation " + self.getOperation())
	}
}

func (self OnlinestoreHandler) empty() error{
	// Get the address key derived from the store address's public key
	AddressKey := self.getNamespaceAddress() + utils.Hex_encryption(self.getAddressFrom())[:constants.TransactionUserAddressLength]
	// recording log
	logger.Info("Got address key " + self.getAddressFrom() + "Address Key " + AddressKey)
	// Get amount from ledger state
	warehouse, err := self.getContext().GetState([]string{AddressKey})
	if err != nil {
		return err
	}
	amount := string(warehouse[AddressKey])
	// GetState() will return empty map if address key doesn't exist in state
	if amount == "" {
		return errors.New("Didn't find the address key associated with " + self.getAddressFrom())
	}
	// Update amount
	updateAmount := 0
	entry := make(map[string][]byte)
	entry[AddressKey] = []byte(strconv.Itoa(updateAmount))
	logger.Info("Empty amount: " + strconv.Itoa(self.getAmount()))
	self.getContext().SetState(entry)
	return nil
}

func (self OnlinestoreHandler) transport() error {

	// get hex code which is in encryption
	addressKeyFrom := self.getNamespaceAddress() + utils.Hex_encryption(self.getAddressFrom())[:constants.TransactionUserAddressLength]
	addressKeyTo := self.getNamespaceAddress() + utils.Hex_encryption(self.getAddressTo())[:constants.TransactionUserAddressLength]

	// Get and validate amount from store state for fromaddress
	warehouse, err := self.getContext().GetState([]string{addressKeyFrom, addressKeyTo})
	if err != nil {
		return err
	}
	amount := string(warehouse[addressKeyFrom])
	if amount == "" {
		return errors.New("Didn't find the address key associated with " + self.getAddressFrom())
	}
	// put amount into int
	FromAmount, err := strconv.Atoi(amount)
	if err != nil {
		return err
	}
	if FromAmount < self.getAmount() {
		// int to string
		return errors.New("Transport amount should be lesser than or equal to " + strconv.Itoa(FromAmount))
	}

	// Get and validate amount from state for toaddress
	amount = string(warehouse[addressKeyTo])
	if amount == "" {
		return errors.New("Didn't find the address key associated with " + self.getAddressTo())
	}
	// get the amount
	Toamount, err := strconv.Atoi(amount)
	if err != nil {
		return err
	}

	// Update amount
	updateAmount := FromAmount - self.getAmount()  // int
	// generate map
	entry := make(map[string][]byte)
	entry[addressKeyFrom] = []byte(strconv.Itoa(updateAmount))
	// recording log
	logger.Info("Debiting amount with " + strconv.Itoa(self.getAmount()))
	// store in state
	self.getContext().SetState(entry)

	// Update amount
	updateAmount = Toamount + self.getAmount()
	entry = make(map[string][]byte)
	entry[addressKeyTo] = []byte(strconv.Itoa(updateAmount))
	logger.Info("Crediting to amount with " + strconv.Itoa(self.getAmount()))
	// store in state
	self.getContext().SetState(entry)
	return nil
}

func (self OnlinestoreHandler) sell() error {
	// Get the address key derived from the address's public key
	addressKey := self.getNamespaceAddress() + utils.Hex_encryption(self.getAddressFrom())[:constants.TransactionUserAddressLength]
	// recording log
	logger.Info("Got address key " + self.getAddressFrom() + "address key " + addressKey)
	// Get amount from state
	warehouse, err := self.getContext().GetState([]string{addressKey})
	if err != nil {
		return err
	}
	amount := string(warehouse[addressKey])
	// GetState() will return empty map if wallet key doesn't exist in state
	if amount == "" {
		return errors.New("Didn't find the address key associated with " + self.getAddressFrom())
	}
	value, err := strconv.Atoi(amount)
	if err != nil {
		return err
	}
	if value < self.getAmount() {
		return errors.New("Sell amount should be lesser than or equal to " + strconv.Itoa(value))
	}
	// Update amount
	updateAmount := value - self.getAmount()
	entry := make(map[string][]byte)
	entry[addressKey] = []byte(strconv.Itoa(updateAmount))
	logger.Info("Selling amount: " + strconv.Itoa(self.getAmount()))
	self.getContext().SetState(entry)
	return nil
}

func (self OnlinestoreHandler) buy() error {
	// Get the address key derived from the address's public key
	addressKey := self.getNamespaceAddress() + utils.Hex_encryption(self.getAddressFrom())[:constants.TransactionUserAddressLength]
	// recording log
	logger.Info("Got address key " + self.getAddressFrom() + "address key " + addressKey)
	// Get amount from state
	warehouse, err := self.getContext().GetState([]string{addressKey})
	if err != nil {
		return err
	}
	amount := string(warehouse[addressKey])
	new_amount := 0
	// GetState() will return empty map if wallet key doesn't exist in state
	if amount == "" {
		logger.Info("This is the first time we buy goods.")
		logger.Info("Creating a new warehouse for the good: " + self.getAddressFrom())
		new_amount = self.getAmount()
	} else {
		var err error
		// get now amount
		new_amount, err = strconv.Atoi(amount)
		// add new amount
		new_amount += self.getAmount()
		if err != nil {
			return err
		}
	}
	// Update amount in state
	entry := make(map[string][]byte)
	entry[addressKey] = []byte(strconv.Itoa(new_amount))
	// recording log
	logger.Info("Buy amount: " + strconv.Itoa(self.getAmount()))
	self.getContext().SetState(entry)
	return nil
}
