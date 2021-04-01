// Copyright (C) 2021 Tweag IO
//
// This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

package config

import (
	"github.com/BurntSushi/toml"
)

type GitStorageConfig struct {
	Remote   string `toml:"remote"`
	Commiter string `toml:"commiter"`
	Email    string `toml:"email"`
}

type NativeStorageConfig struct {
}

type StorageConfig struct {
	Type   string               `toml:"type"`
	Git    *GitStorageConfig    `toml:"git"`
	Native *NativeStorageConfig `toml:"native"`
}

type GRPCTransportConfig struct {
	Remote string `toml:"remote"`
}

type TransportConfig struct {
	Type string               `toml:"type"`
	GRPC *GRPCTransportConfig `toml:"grpc"`
}

type ED25519SignerConfig struct {
	PrivateKeyPath string `toml:"private-key-path"`
}

type SignerConfig struct {
	Type      string               `toml:"type"`
	KeyType   string               `toml:"key-type"`
	PublicKey string               `toml:"public-key"`
	ED25519   *ED25519SignerConfig `toml:"ed25519"`
}

type LogConfig struct {
	Name      string            `toml:"name"`
	Mode      string            `toml:"mode"`
	Storage   *StorageConfig    `toml:"storage"`
	Transport *TransportConfig  `toml:"transport"`
	Signer    *SignerConfig     `toml:"signer"`
	Meta      map[string]string `toml:"meta"`
}

type LuaDeciderConfig struct {
	Script string `toml:"script"`
}

type LogNameDeciderConfig struct {
	Name string `toml:"name"`
}

type PercentageDeciderConfig struct {
	Minimum int `toml:"minimum"`
}

type DeciderConfig struct {
	Engine     string                   `toml:"engine"`
	Lua        *LuaDeciderConfig        `toml:"lua"`
	LogName    *LogNameDeciderConfig    `toml:"logname"`
	Percentage *PercentageDeciderConfig `toml:"percentage"`
}

type Config struct {
	Deciders []*DeciderConfig `toml:"decider"`
	Logs     []*LogConfig     `toml:"log"`
}

func NewConfigFromFile(path string) (*Config, error) {
	conf := &Config{}

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}
