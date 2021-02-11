package chopsticks

import (
	"bytes"
	"encoding/json"
	"errors"
)

// BEGIN manifest (app.json/manifest.json)

// you could find the manifest file schema at ./schema/manifest.json

// generated from JSON Schema using quicktype, but with modification
// To parse and unparse this JSON data, add this code to your project and do:
//
//    manifest, err := UnmarshalManifest(bytes)
//    bytes, err = manifest.Marshal()

func UnmarshalManifest(data []byte) (Manifest, error) {
	var r Manifest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Manifest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Manifest) ToPrettyJson() (string, error) {
	p, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}

	return string(p), nil
}

type Manifest struct {
	Empty        *StringOrArrayOfStrings                          `json:"##,omitempty"` // A comment.
	Version      string                                           `json:"version,omitempty"`
	Homepage     *string                                          `json:"homepage,omitempty"`
	Description  *string                                          `json:"description,omitempty"`
	License      *LicenseUnion                                    `json:"license,omitempty"`
	URL          *StringOrArrayOfStrings                          `json:"url,omitempty"`
	Cookie       map[string]interface{}                           `json:"cookie,omitempty"` // Undocumented: Found at https://github.com/se35710/scoop-java/search?l=JSON&q=cookie
	Hash         *Hash                                            `json:"hash,omitempty"`
	Architecture *ManifestArchitecture                            `json:"architecture,omitempty"`
	ExtractDir   *StringOrArrayOfStrings                          `json:"extract_dir,omitempty"`
	ExtractTo    *StringOrArrayOfStrings                          `json:"extract_to,omitempty"`
	Innosetup    *bool                                            `json:"innosetup,omitempty"` // True if the installer InnoSetup based. Found in; https://github.com/lukesampson/scoop/search?l=JSON&q=innosetup
	Installer    *ManifestInstaller                               `json:"installer,omitempty"`
	Uninstaller  *Uninstaller                                     `json:"uninstaller,omitempty"`
	PreInstall   *StringOrArrayOfStrings                          `json:"pre_install,omitempty"`
	PostInstall  *StringOrArrayOfStrings                          `json:"post_install,omitempty"`
	EnvAddPath   *StringOrArrayOfStrings                          `json:"env_add_path,omitempty"`
	EnvSet       map[string]interface{}                           `json:"env_set,omitempty"`
	Psmodule     *ManifestPsmodule                                `json:"psmodule,omitempty"`
	Bin          *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"bin,omitempty"`
	Shortcuts    [][]string                                       `json:"shortcuts,omitempty"`
	Persist      *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"persist,omitempty"`
	Checkver     *Checkver                                        `json:"checkver,omitempty"`
	Autoupdate   *Autoupdate                                      `json:"autoupdate,omitempty"`
	Depends      *StringOrArrayOfStrings                          `json:"depends,omitempty"`
	Suggest      *Suggest                                         `json:"suggest,omitempty"`
	Notes        *StringOrArrayOfStrings                          `json:"notes,omitempty"`
}

type ManifestArchitecture struct {
	The32Bit *The32BitClass `json:"32bit,omitempty"`
	The64Bit *The32BitClass `json:"64bit,omitempty"`
}

type The32BitClass struct {
	URL         *StringOrArrayOfStrings                          `json:"url,omitempty"`
	Hash        *Hash                                            `json:"hash,omitempty"`
	ExtractDir  *StringOrArrayOfStrings                          `json:"extract_dir,omitempty"`
	ExtractTo   *StringOrArrayOfStrings                          `json:"extract_to,omitempty"`
	Installer   *ManifestInstaller                               `json:"installer,omitempty"`
	Uninstaller *Uninstaller                                     `json:"uninstaller,omitempty"`
	PreInstall  *StringOrArrayOfStrings                          `json:"pre_install,omitempty"`
	PostInstall *StringOrArrayOfStrings                          `json:"post_install,omitempty"`
	EnvAddPath  *StringOrArrayOfStrings                          `json:"env_add_path,omitempty"`
	EnvSet      map[string]interface{}                           `json:"env_set,omitempty"`
	Bin         *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"bin,omitempty"`
	Shortcuts   [][]string                                       `json:"shortcuts,omitempty"`
	Checkver    *Checkver                                        `json:"checkver,omitempty"`
}

type CheckverClass struct {
	Github    *string `json:"github,omitempty"`
	Jp        *string `json:"jp,omitempty"` // Same as 'jsonpath'
	Jsonpath  *string `json:"jsonpath,omitempty"`
	Re        *string `json:"re,omitempty"` // Same as 'regex'
	Regex     *string `json:"regex,omitempty"`
	Replace   *string `json:"replace,omitempty"` // Allows rearrange the regexp matches
	Reverse   *bool   `json:"reverse,omitempty"` // Reverse the order of regex matches
	URL       *string `json:"url,omitempty"`
	Useragent *string `json:"useragent,omitempty"`
	Xpath     *string `json:"xpath,omitempty"`
}

type ManifestInstaller struct {
	Args   *StringOrArrayOfStrings `json:"args,omitempty"`
	File   *string                 `json:"file,omitempty"`
	Keep   *bool                   `json:"keep,omitempty"`
	Script *StringOrArrayOfStrings `json:"script,omitempty"`
}

type Uninstaller struct {
	Args   *StringOrArrayOfStrings `json:"args,omitempty"`
	File   *string                 `json:"file,omitempty"`
	Script *StringOrArrayOfStrings `json:"script,omitempty"`
}

type Autoupdate struct {
	URL          *StringOrArrayOfStrings                          `json:"url,omitempty"`
	Hash         *HashExtractionOrArrayOfHashExtractions          `json:"hash,omitempty"`
	ExtractDir   *StringOrArrayOfStrings                          `json:"extract_dir,omitempty"`
	ExtractTo    *StringOrArrayOfStrings                          `json:"extract_to,omitempty"`
	Installer    *AutoupdateInstaller                             `json:"installer,omitempty"`
	PreInstall   *StringOrArrayOfStrings                          `json:"pre_install,omitempty"`
	PostInstall  *StringOrArrayOfStrings                          `json:"post_install,omitempty"`
	EnvAddPath   *StringOrArrayOfStrings                          `json:"env_add_path,omitempty"`
	EnvSet       map[string]interface{}                           `json:"env_set,omitempty"`
	Psmodule     *The32BitPsmodule                                `json:"psmodule,omitempty"`
	Bin          *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"bin,omitempty"`
	Shortcuts    [][]string                                       `json:"shortcuts,omitempty"`
	Persist      *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"persist,omitempty"`
	Architecture *AutoupdateArchitecture                          `json:"architecture,omitempty"`
	Note         *StringOrArrayOfStrings                          `json:"note,omitempty"`
}

type AutoupdateArchitecture struct {
	The32Bit *AutoupdateArch `json:"32bit,omitempty"`
	The64Bit *AutoupdateArch `json:"64bit,omitempty"`
}

type AutoupdateArch struct {
	URL         *StringOrArrayOfStrings                          `json:"url,omitempty"`
	Hash        *HashExtractionOrArrayOfHashExtractions          `json:"hash,omitempty"`
	ExtractDir  *StringOrArrayOfStrings                          `json:"extract_dir,omitempty"`
	ExtractTo   *StringOrArrayOfStrings                          `json:"extract_to,omitempty"`
	Installer   *AutoupdateInstaller                             `json:"installer,omitempty"`
	PreInstall  *StringOrArrayOfStrings                          `json:"pre_install,omitempty"`
	PostInstall *StringOrArrayOfStrings                          `json:"post_install,omitempty"`
	EnvAddPath  *StringOrArrayOfStrings                          `json:"env_add_path,omitempty"`
	EnvSet      map[string]interface{}                           `json:"env_set,omitempty"`
	Psmodule    *The32BitPsmodule                                `json:"psmodule,omitempty"`
	Bin         *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"bin,omitempty"`
	Shortcuts   [][]string                                       `json:"shortcuts,omitempty"`
	Persist     *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings `json:"persist,omitempty"`
	Note        *StringOrArrayOfStrings                          `json:"note,omitempty"`
}

type HashExtraction struct {
	Find     *string `json:"find,omitempty"` // Same as 'regex'
	Jp       *string `json:"jp,omitempty"`   // Same as 'jsonpath'
	Jsonpath *string `json:"jsonpath,omitempty"`
	Mode     *Mode   `json:"mode,omitempty"`
	Regex    *string `json:"regex,omitempty"`
	URL      *string `json:"url,omitempty"`
	Xpath    *string `json:"xpath,omitempty"`
}

type AutoupdateInstaller struct {
	File *string `json:"file,omitempty"`
}

type The32BitPsmodule struct {
	Name *string `json:"name,omitempty"`
}

type LicenseClass struct {
	Identifier string  `json:"identifier,omitempty"`
	URL        *string `json:"url,omitempty"`
}

type ManifestPsmodule struct {
	Name *string `json:"name,omitempty"`
}

type Suggest struct {
}

type Mode string

const (
	Download    Mode = "download"
	Extract     Mode = "extract"
	Fosshub     Mode = "fosshub"
	JSON        Mode = "json"
	Metalink    Mode = "metalink"
	RDF         Mode = "rdf"
	Sourceforge Mode = "sourceforge"
	Xpath       Mode = "xpath"
)

type StringOrArrayOfStringsOrAnArrayOfArrayOfStrings struct {
	String     *string
	UnionArray []StringOrArrayOfStrings
}

func (x *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StringOrArrayOfStringsOrAnArrayOfArrayOfStrings) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}

// A comment.
type StringOrArrayOfStrings struct {
	String      *string
	StringArray []string
}

func (x *StringOrArrayOfStrings) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StringOrArrayOfStrings) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type Checkver struct {
	CheckverClass *CheckverClass
	String        *string
}

func (x *Checkver) UnmarshalJSON(data []byte) error {
	x.CheckverClass = nil
	var c CheckverClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CheckverClass = &c
	}
	return nil
}

func (x *Checkver) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.CheckverClass != nil, x.CheckverClass, false, nil, false, nil, false)
}

type Hash struct {
	String      *string
	StringArray []string
}

func (x *Hash) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Hash) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type HashExtractionOrArrayOfHashExtractions struct {
	HashExtraction      *HashExtraction
	HashExtractionArray []HashExtraction
}

func (x *HashExtractionOrArrayOfHashExtractions) UnmarshalJSON(data []byte) error {
	x.HashExtractionArray = nil
	x.HashExtraction = nil
	var c HashExtraction
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.HashExtractionArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HashExtraction = &c
	}
	return nil
}

func (x *HashExtractionOrArrayOfHashExtractions) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.HashExtractionArray != nil, x.HashExtractionArray, x.HashExtraction != nil, x.HashExtraction, false, nil, false, nil, false)
}

type LicenseUnion struct {
	LicenseClass *LicenseClass
	String       *string
}

func (x *LicenseUnion) UnmarshalJSON(data []byte) error {
	x.LicenseClass = nil
	var c LicenseClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.LicenseClass = &c
	}
	return nil
}

func (x *LicenseUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.LicenseClass != nil, x.LicenseClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}

// END manifest

// BEGIN install info (install.json)
// example install.json file:
// {
//    // url OR bucket
//    "url": "C:\\path\\to\\app\\manifest.json",
//    "bucket": "bucket-name",
//    // available arches: 32bit, 64bit
//    "architecture": "64bit",
//   // bool value, is this app held
//    "hold": false
// }

// This file was generated from JSON Schema using quicktype.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    installInfo, err := UnmarshalInstallInfo(bytes)
//    bytes, err = installInfo.Marshal()

// UnmarshalInstallInfo 将字节流形式的 `install.json` 转换成 InstallInfo 结构体
func UnmarshalInstallInfo(data []byte) (InstallInfo, error) {
	var r InstallInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal 将 InstallInfo 结构体转换为字节流
func (r *InstallInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ToPrettyJson 将 InstallInfo 信息以 pretty-print json 形式输出
func (r *InstallInfo) ToPrettyJson() (string, error) {
	p, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}

	return string(p), nil
}

// InstallInfo 是安装好后，存在于 apps/version/ 下 ，描述当前应用，拥有以下字段：
// 从何处安装（url, bucket）
// 架构信息（architecture）
// 是否版本冻结（hold）
type InstallInfo struct {
	URL          string `json:"url,omitempty"`
	Bucket       string `json:"bucket,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Hold         bool   `json:"hold,omitempty"`
}

// END install info
