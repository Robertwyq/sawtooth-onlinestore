/**
 * Author：Robert WYQ
 * Date：2019-7-22 ——  2019-8-31
 * Description: To create an onlinestore APP base on the sawtooth1.1.5
 * Level: main function of processor
 * ------------------------------------------------------------------------------
 */

package main

import (
	"fmt"
	"github.com/hyperledger/sawtooth-sdk-go/logging"
	myprocessor "github.com/hyperledger/sawtooth-sdk-go/processor"
	"github.com/jessevdk/go-flags"
	"os"
	myhandler "./handlers"
	"syscall"
)
// Copied sample transaction processor handler from sdk-example
// handler：异步消息的处理机制
type Opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Increase verbosity"`
	Connect string `short:"C" long:"connect" description:"Validator component endpoint to connect to" default:"tcp://localhost:4004"`
}
// 基础版的就只有Verbose和Connect，intkey、xo，复杂的类似small bank有Queue和Threads
// Queue：消息的队列？最大缓存消息数？
// Threads：选择并行执行的线程数

func main() {

	var opts Opts

	// create a new log
	logger := logging.Get()

	// create a parser
	parser := flags.NewParser(&opts, flags.Default)

	remaining, err := parser.Parse()

	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			logger.Errorf("Failed to parse args: %v", err)
			os.Exit(2)
		}
	}

	if len(remaining) > 0 {
		fmt.Printf("Error: Unrecognized arguments passed: %v\n", remaining)
		os.Exit(2)
	}

	// create endpoint for connecting
	endpoint := opts.Connect

	// define different log level
	switch len(opts.Verbose) {
	case 2:
		logger.SetLevel(logging.DEBUG)
	case 1:
		logger.SetLevel(logging.INFO)
	default:
		logger.SetLevel(logging.WARN)
	}

	logger.Debugf("command line arguments: %v", os.Args)
	logger.Debugf("verbose = %v\n", len(opts.Verbose))
	logger.Debugf("endpoint = %v\n", endpoint)

	// create a handler in our defined struct
	handler := &myhandler.OnlinestoreHandler{}

	// create a processor link to the endpoint
	processor := myprocessor.NewTransactionProcessor(endpoint)

	// AddHandler adds the given handler to the TransactionProcessor so it can receive transaction processing requests.
	// All handlers must be added prior to starting the processor.
	processor.AddHandler(handler)

	// ShutdownOnSignal sets up signal handling to shutdown the processor when one of the signals passed is received.
	processor.ShutdownOnSignal(syscall.SIGINT, syscall.SIGTERM)

	// Start connects the TransactionProcessor to a validator and starts listening for requests and routing them to an appropriate handler.
	err = processor.Start()
	// recording error if connecting error
	if err != nil {
		logger.Error("Processor stopped: ", err)
	}
}