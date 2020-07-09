package infraview_test

import (
	"io/ioutil"
	"testing"

	"github.com/cycloidio/infraview/graph"
	"github.com/cycloidio/infraview/infraview"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromState_AWS(t *testing.T) {
	t.Run("SuccessSG", func(t *testing.T) {
		src, err := ioutil.ReadFile("./testdata/aws_sg.json")
		require.NoError(t, err)

		g, cfg, err := infraview.FromState(src)
		require.NoError(t, err)
		require.NotNil(t, g)
		require.NotNil(t, cfg)

		eg := &graph.Graph{
			Nodes: []*graph.Node{
				&graph.Node{
					Canonical: "aws_lb.tQBgz",
				},
				&graph.Node{
					Canonical: "aws_launch_template.vIkyE",
				},
				&graph.Node{
					Canonical: "aws_db_instance.Cpbzf",
				},
			},
			Edges: []*graph.Edge{
				&graph.Edge{
					Source:     "aws_lb.tQBgz",
					Target:     "aws_launch_template.vIkyE",
					Canonicals: []string{"aws_security_group.rZnGI", "aws_security_group.YPHPR"},
				},
				&graph.Edge{
					Source:     "aws_launch_template.vIkyE",
					Target:     "aws_db_instance.Cpbzf",
					Canonicals: []string{"aws_security_group.YPHPR", "aws_security_group.LHwFh"},
				},
			},
		}

		assertEqualGraph(t, eg, g, cfg)
	})

	t.Run("SuccessSGR", func(t *testing.T) {
		src, err := ioutil.ReadFile("./testdata/aws_sgr.json")
		require.NoError(t, err)

		g, cfg, err := infraview.FromState(src)
		require.NoError(t, err)
		require.NotNil(t, g)
		require.NotNil(t, cfg)

		eg := &graph.Graph{
			Nodes: []*graph.Node{
				&graph.Node{
					Canonical: "aws_elb.tMVdH",
				},
				&graph.Node{
					Canonical: "aws_instance.TObJL",
				},
				&graph.Node{
					Canonical: "aws_db_instance.qktIK",
				},
				&graph.Node{
					Canonical: "aws_elasticache_cluster.VUhMF",
				},
			},
			Edges: []*graph.Edge{
				&graph.Edge{
					Source:     "aws_elb.tMVdH",
					Target:     "aws_instance.TObJL",
					Canonicals: []string{"aws_security_group.kuDkz", "aws_security_group_rule.pMOSN", "aws_security_group.UKblk"},
				},
				&graph.Edge{
					Source:     "aws_instance.TObJL",
					Target:     "aws_db_instance.qktIK",
					Canonicals: []string{"aws_security_group.mzSGd", "aws_security_group.kuDkz"},
				},
				&graph.Edge{
					Source:     "aws_instance.TObJL",
					Target:     "aws_elasticache_cluster.VUhMF",
					Canonicals: []string{"aws_security_group.KaWAd", "aws_security_group.kuDkz"},
				},
			},
		}

		assertEqualGraph(t, eg, g, cfg)
	})

	t.Run("WithCount", func(t *testing.T) {
		src, err := ioutil.ReadFile("./testdata/aws_with_count.json")
		require.NoError(t, err)

		g, cfg, err := infraview.FromState(src)
		require.NoError(t, err)
		require.NotNil(t, g)
		require.NotNil(t, cfg)
		assert.Len(t, g.Nodes, 2)
	})
}

func TestFromState_OpenStack(t *testing.T) {
	t.Run("SuccessLB", func(t *testing.T) {
		src, err := ioutil.ReadFile("./testdata/openstack_lb.json")
		require.NoError(t, err)

		g, cfg, err := infraview.FromState(src)
		require.NoError(t, err)
		require.NotNil(t, g)
		require.NotNil(t, cfg)

		eg := &graph.Graph{
			Nodes: []*graph.Node{
				&graph.Node{
					Canonical: "openstack_compute_instance_v2.AaCFA",
				},
				&graph.Node{
					Canonical: "openstack_compute_instance_v2.gZfYc",
				},
				&graph.Node{
					Canonical: "openstack_lb_loadbalancer_v2.PPdjL",
				},
			},
			Edges: []*graph.Edge{
				&graph.Edge{
					Target: "openstack_compute_instance_v2.gZfYc",
					Source: "openstack_compute_instance_v2.AaCFA",
					Canonicals: []string{
						"openstack_networking_port_v2.GQStv",
						"openstack_networking_port_v2.PimKo",
						"openstack_networking_secgroup_rule_v2.uzQon",
						"openstack_networking_secgroup_v2.KFnza",
						"openstack_networking_secgroup_v2.ilWCI",
					},
				},
				&graph.Edge{
					Target: "openstack_compute_instance_v2.gZfYc",
					Source: "openstack_lb_loadbalancer_v2.PPdjL",
					Canonicals: []string{
						"openstack_lb_listener_v2.lzSKa",
						"openstack_lb_member_v2.FbSow",
						"openstack_lb_pool_v2.nwYyz",
					},
				},
			},
		}

		assertEqualGraph(t, eg, g, cfg)
	})
	t.Run("SuccessSG", func(t *testing.T) {
		src, err := ioutil.ReadFile("./testdata/openstack_sg.json")
		require.NoError(t, err)

		g, cfg, err := infraview.FromState(src)
		require.NoError(t, err)
		require.NotNil(t, g)
		require.NotNil(t, cfg)

		eg := &graph.Graph{
			Nodes: []*graph.Node{
				&graph.Node{
					Canonical: "openstack_compute_instance_v2.sjSbA",
				},
				&graph.Node{
					Canonical: "openstack_compute_instance_v2.PiGtZ",
				},
			},
			Edges: []*graph.Edge{
				&graph.Edge{
					Target: "openstack_compute_instance_v2.sjSbA",
					Source: "openstack_compute_instance_v2.PiGtZ",
					Canonicals: []string{
						"openstack_networking_port_v2.QWPcW",
						"openstack_networking_port_v2.tZVzk",
						"openstack_networking_secgroup_rule_v2.QdNte",
						"openstack_networking_secgroup_v2.DKQQX",
						"openstack_networking_secgroup_v2.ievUc",
					},
				},
			},
		}

		assertEqualGraph(t, eg, g, cfg)
	})

}

func TestFromState_FlexibelEngine(t *testing.T) {
}
