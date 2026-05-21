package vmware

import (
	"fmt"
	"log"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25/types"
)

func (vc *Vcenter) GetWmksTicket(vmName string) (string, error) {
	log.Println("Generando token VMKS para ", vmName)

	finder := find.NewFinder(vc.Client.Client, true)
	dc, err := finder.DefaultDatacenter(vc.Ctx)
	if err != nil {
		return "", err
	}
	finder.SetDatacenter(dc)

	vm, err := finder.VirtualMachine(vc.Ctx, vmName)
	if err != nil {
		return "", err
	}

	mksTicket, err := vm.AcquireTicket(vc.Ctx, "webmks")
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("wss://%s/ticket/%s", mksTicket.Host, mksTicket.Ticket)
	log.Println(url)
	return url, nil
}

func (vc *Vcenter) GetWmksParams(vmName string) (*types.VirtualMachineTicket, error) {
	log.Println("Generando token VMKS para ", vmName)

	finder := find.NewFinder(vc.Client.Client, true)
	dc, err := finder.DefaultDatacenter(vc.Ctx)
	if err != nil {
		return nil, err
	}
	finder.SetDatacenter(dc)

	vm, err := finder.VirtualMachine(vc.Ctx, vmName)
	if err != nil {
		return nil, err
	}

	mksTicket, err := vm.AcquireTicket(vc.Ctx, "webmks")
	if err != nil {
		return nil, err
	}

	return mksTicket, nil
}
