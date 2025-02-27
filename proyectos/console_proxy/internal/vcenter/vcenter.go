package vmware

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/vmware/govmomi"
)

type Vcenter struct {
	URL    string
	Client *govmomi.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

var (
	vcuser = url.QueryEscape(os.Getenv("vc_user"))
	vcpass = url.QueryEscape(os.Getenv("vc_pass"))
)

var vcDomains = map[string]string{
	"adresvc.ifx-adr.local":    "ifx-adr.local",
	"hipovc.ifx-hip.local":     "ifx-hip.local",
	"tovavc.ifx-tva.local":     "ifx-tva.local",
	"vc70.vmng.lan":            "vsphere.local",
	"vcbaq.ifxcorp.com":        "ifx.local",
	"vcbognp.ifxcorp.com":      "ifxnp.local",
	"vcbsq.ifxcorp.com":        "ifx.local",
	"vcbue.ifxcorp.com":        "ifx.local",
	"vcdtc.ifxcorp.com":        "ifx.local",
	"vcdtcmgmt.ifxcorp.com":    "ifx.local",
	"vcgua.ifxcorp.com":        "vsphere.local",
	"vcigt.ifx-igt.local":      "ifx-igt.local",
	"vclim.ifxcorp.com":        "ifx.local",
	"vcmgn.ifxcorp.com":        "ifx.local",
	"vcmia.ifxcorp.com":        "ifx.local",
	"vcmiamgmt.ifxcorp.com":    "ifx.local",
	"vcpan.ifxcorp.com":        "ifx.local",
	"vcqua.quantic.local":      "quantic.local",
	"vcsap.ifxcorp.com":        "ifx.local",
	"vctst.ifxcorp.com":        "ifx.local",
	"vcwbp.ifxcorp.com":        "ifx.local",
	"wlufinetvc.ifx-ufi.local": "ifx-ufi.local",
}

func NewVcenter(vcenter string) (Vcenter, error) {
	ctx, cancel := context.WithCancel(context.Background())
	vc := Vcenter{Ctx: ctx, Cancel: cancel}
	userDomain, exist := vcDomains[vcenter]
	if !exist {
		return vc, fmt.Errorf("no se encuentra dominio para el vcenter %s", vcenter)
	}
	vciuser := strings.Replace(vcuser, "ifx.local", userDomain, 1)

	vcenterURL, err := url.Parse(fmt.Sprintf("https://%s:%s@%s/sdk", vciuser, vcpass, vcenter))
	if err != nil {
		return vc, err
	}

	client, err := govmomi.NewClient(ctx, vcenterURL, true)
	if err != nil {
		return vc, err
	}

	vc.Client = client
	vc.URL = vcenter
	log.Println("Sesión iniciana en:", vcenter)
	return vc, nil
}

func (vc *Vcenter) Logout() {
	if vc.Client != nil {
		_ = vc.Client.Logout(vc.Ctx)
	}
	if vc.Cancel != nil {
		vc.Cancel()
	}
}
