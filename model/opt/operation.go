package opt

type Operation Opt

func (o Operation) IsSn(sn string) bool {
	var is = false
	for i := 0; i < len(o.Data); i++ {
		if sn == o.Data[i].Sn {
			is = true
			break
		}
		continue
	}
	return is
}

func(o Operation) SnInMode(sn string) string {
	var mode string

	for i := 0; i < len(o.Data); i++ {

		if sn == o.Data[i].Sn {
			mode = o.Data[i].Platform
			break
		}
		continue
	}

	if mode == "nuxes" {
		mode = "nexus"
	}else if mode =="nuxes-watsons" {
		mode = "watsons"
	} else if mode =="watsons-ha" {
		mode = "watsonsha"
	}

	return mode
}
