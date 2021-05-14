package canary

import "github.com/buildpacks/libcnb"

type Detect struct{}

func (d Detect) Detect(context libcnb.DetectContext) (libcnb.DetectResult, error) {
	return libcnb.DetectResult{
		Pass: true,
		Plans: []libcnb.BuildPlan{
			{
				Provides: []libcnb.BuildPlanProvide{
					{Name: "canary"},
				},
				Requires: []libcnb.BuildPlanRequire{
					{Name: "canary"},
				},
			},
		},
	}, nil
}
