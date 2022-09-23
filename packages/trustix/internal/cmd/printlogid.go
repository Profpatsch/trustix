// Copyright (C) 2021 Tweag IO
//
// This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"log"
	"github.com/nix-community/trustix/packages/trustix/internal/protocols"
	"github.com/nix-community/trustix/packages/trustix-proto/api"
	conf "github.com/nix-community/trustix/packages/trustix/internal/config"

	"github.com/spf13/cobra"
)

var printLogIdConfigPath string
var printLogIdSigner string
var printLogIdProtocol string
var printLogIdPublicKeyType string
var printLogIdPublicKey string
var printLogIdMode string

var modeDefault = api.Log_LogModes_name[api.Log_LogModes(0)]

var printLogIdCmd = &cobra.Command{
	Use:   "print-log-id",
	Short: "Print the log-id from a publisher",
	RunE: func(cmd *cobra.Command, args []string) error {

		switch printLogIdPublicKeyType {
		case "ed25519":
		default:
			log.Fatalf("Unhandled key type: %s", printLogIdPublicKeyType)
		}

		var mode, ok = api.Log_LogModes_value[printLogIdMode]
		if !ok {
			log.Fatalf("Unrecognized mode: %s", printLogIdMode)
		}

    if printLogIdConfigPath != "" {
      config, err := conf.NewConfigFromFile(printLogIdConfigPath)
      if err != nil {
        log.Fatal(err)
      }

			var found = false

			for _, publisherConfig := range config.Publishers {
				if printLogIdPublicKeyType == publisherConfig.PublicKey.Type &&
					(printLogIdPublicKey == "" || printLogIdPublicKey == publisherConfig.PublicKey.Pub) &&
					(printLogIdSigner == ""    || printLogIdSigner    == publisherConfig.Signer) &&
					(printLogIdProtocol == ""  || printLogIdProtocol  == publisherConfig.Protocol) {
					if found {
						log.Fatal("More than one publisher matches the criteria given.")
					} else {
						found = true
						printLogIdPublicKey = publisherConfig.PublicKey.Pub
						printLogIdSigner    = publisherConfig.Signer
						printLogIdProtocol  = publisherConfig.Protocol
					}
				}
			}

			if !found {
				log.Fatal("Could not find a log that matches all criteria specified.")
			}
		} else if printLogIdProtocol == "" || printLogIdPublicKey == "" || printLogIdPublicKeyType == "" || printLogIdMode == "" {
			log.Fatal("You must either specify a config path, or specify all necessary log settings via command line flags.")
		}

    protocol, err := protocols.Get(printLogIdProtocol)
		if err != nil {
			log.Fatal(err)
		}


    pubKey := conf.PublicKey{Type: printLogIdPublicKeyType, Pub: printLogIdPublicKey}
    pubBytes, err := pubKey.Decode()
    if err != nil {
      log.Fatal(err)
    }

		logID = protocol.LogID(printLogIdPublicKeyType, pubBytes, mode)

		fmt.Println(logID)

		return nil
	},
}

func initPrintLogId() {
	printLogIdCmd.Flags().StringVar(&printLogIdPublicKeyType, "pubkey-type", "ed25519", "Type of public key")
	printLogIdCmd.Flags().StringVar(&printLogIdPublicKey, "pubkey", "", "Public key of the log")
	printLogIdCmd.Flags().StringVar(&printLogIdConfigPath, "config", "", "Configuration that contains the log")
	printLogIdCmd.Flags().StringVar(&printLogIdSigner, "signer", "", "The log's signer (only used for identifying the log in the config)")
	printLogIdCmd.Flags().StringVar(&printLogIdProtocol, "protocol", "", "The protcol of the log")
	printLogIdCmd.Flags().StringVar(&printLogIdMode, "mode", modeDefault, "") // TODO [TB]: What is "mode"? Add a description
}
