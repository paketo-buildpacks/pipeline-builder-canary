package canary_test

import (
	"github.com/paketo-buildpacks/pipeline-builder-canary/canary"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		build canary.Build
		ctx   libcnb.BuildContext
	)

	it("does nothing without plan", func() {
		Expect(build.Build(ctx)).To(Equal(libcnb.BuildResult{}))
	})

	it("adds metadata to result", func() {
		ctx.Plan = libcnb.BuildpackPlan{
			Entries: []libcnb.BuildpackPlanEntry{
				{
					Name: "canary",
				},
			},
		}

		Expect(build.Build(ctx)).To(Equal(libcnb.NewBuildResult()))
	})

}
