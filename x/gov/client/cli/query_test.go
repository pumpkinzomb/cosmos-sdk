package cli_test

import (
	"fmt"
	"strings"

	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
)

func (s *CLITestSuite) TestCmdParams() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"json output",
			[]string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			"--output=json",
		},
		{
			"text output",
			[]string{},
			"",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryParams()
			cmd.SetArgs(tc.args)

			if len(tc.args) != 0 {
				s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
			}
		})
	}
}

func (s *CLITestSuite) TestCmdParam() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"voting params",
			[]string{
				"voting",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			`voting --output=json`,
		},
		{
			"tally params",
			[]string{
				"tallying",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			`tallying --output=json`,
		},
		{
			"deposit params",
			[]string{
				"deposit",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			`deposit --output=json`,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryParam()
			cmd.SetArgs(tc.args)
			s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
		})
	}
}

func (s *CLITestSuite) TestCmdProposer() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"without proposal id",
			[]string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"--output=json",
		},
		{
			"json output",
			[]string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"1 --output=json",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryProposer()
			cmd.SetArgs(tc.args)
			s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
		})
	}
}

func (s *CLITestSuite) TestCmdTally() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"without proposal id",
			[]string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"--output=json",
		},
		{
			"json output",
			[]string{
				"2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"2 --output=json",
		},
		{
			"json output",
			[]string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"1 --output=json",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryTally()
			cmd.SetArgs(tc.args)
			s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
		})
	}
}

func (s *CLITestSuite) TestCmdGetProposal() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"get non existing proposal",
			[]string{
				"10",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"10 --output=json",
		},
		{
			"get proposal with json response",
			[]string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"1 --output=json",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryProposal()
			cmd.SetArgs(tc.args)
			s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
		})
	}
}

func (s *CLITestSuite) TestCmdGetProposals() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"get proposals as json response",
			[]string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"--output=json",
		},
		{
			"get proposals with invalid status",
			[]string{
				"--status=unknown",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"--status=unknown --output=json",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryProposals()
			cmd.SetArgs(tc.args)
			s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
		})
	}
}

func (s *CLITestSuite) TestCmdQueryDeposits() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"get deposits of non existing proposal",
			[]string{
				"10",
			},
			"10",
		},
		{
			"get deposits(valid req)",
			[]string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"1 --output=json",
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryDeposits()
			cmd.SetArgs(tc.args)

			if len(tc.args) != 0 {
				s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
			}
		})
	}
}

func (s *CLITestSuite) TestCmdQueryDeposit() {
	val := testutil.CreateKeyringAccounts(s.T(), s.kr, 1)

	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"get deposit with no depositer",
			[]string{
				"1",
			},
			"1",
		},
		{
			"get deposit with wrong deposit address",
			[]string{
				"1",
				"wrong address",
			},
			"1 wrong address",
		},
		{
			"get deposit (valid req)",
			[]string{
				"1",
				val[0].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			fmt.Sprintf("1 %s --output=json", val[0].Address.String()),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryDeposit()
			cmd.SetArgs(tc.args)

			if len(tc.args) != 0 {
				s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
			}
		})
	}
}

func (s *CLITestSuite) TestCmdQueryVotes() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"get votes with no proposal id",
			[]string{},
			"true",
		},
		{
			"get votes of non existed proposal",
			[]string{
				"10",
			},
			"10",
		},
		{
			"vote for invalid proposal",
			[]string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			"1 --output=json",
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryVotes()
			cmd.SetArgs(tc.args)

			if len(tc.args) != 0 {
				s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
			}
		})
	}
}

func (s *CLITestSuite) TestCmdQueryVote() {
	val := testutil.CreateKeyringAccounts(s.T(), s.kr, 1)

	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"get vote of non existing proposal",
			[]string{
				"10",
				val[0].Address.String(),
			},
			fmt.Sprintf("10 %s", val[0].Address.String()),
		},
		{
			"get vote by wrong voter",
			[]string{
				"1",
				"wrong address",
			},
			"1 wrong address",
		},
		{
			"vote for valid proposal",
			[]string{
				"1",
				val[0].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			fmt.Sprintf("1 %s --output=json", val[0].Address.String()),
		},
		{
			"split vote for valid proposal",
			[]string{
				"3",
				val[0].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			fmt.Sprintf("3 %s --output=json", val[0].Address.String()),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryVote()
			cmd.SetArgs(tc.args)

			if len(tc.args) != 0 {
				s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
			}
		})
	}
}
