package cli

import "testing"

// Note: add -v to any of these commands to enable verbose logging

func TestAssets(t *testing.T) {
	cli, _ := newTestCLI()
	cli.TestCommand("ns test")

	expectOutput(t, cli, "error", "asset set USD someissuer")
	expectOutput(t, cli, "error", "asset set USD SBWP26IQVZIH52ZCBW4ETX4I4XJZZHNTW5PNWNKSMM25WRBKTJQ7DWGD")
	expectOutput(t, cli, "", "asset set USD GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM")

	cli.TestCommand("account new chase_bank")
	expectOutput(t, cli, "", "asset set USD chase_bank")

	expectOutput(t, cli, "USD", "asset code USD")
	expectOutput(t, cli, "credit_alphanum4", "asset type USD")

	expectOutput(t, cli, "", "asset set USD-chase GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM --code USD")
	expectOutput(t, cli, "USD", "asset code USD-chase")
	expectOutput(t, cli, "GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM", "asset issuer USD-chase")

	expectOutput(t, cli, "error", "asset set USD-bad GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM --type credit16")
	expectOutput(t, cli, "", "asset set USD-bad GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM --type native")

	cli.TestCommand("asset del USD-chase")
	expectOutput(t, cli, "error", "asset issuer USD-chase")

	expectOutput(t, cli, "GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM", "asset issuer USD:GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM")
	cli.TestCommand("account set citibank GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM")
	expectOutput(t, cli, "GBY7XDYKXBDHQ2B523SF7K6BNJNRYHVQMWY7AYAEKTYLCQMYVFHL57UM", "asset issuer USD:citibank")
	expectOutput(t, cli, "USD", "asset code USD:citibank")
	expectOutput(t, cli, "credit_alphanum4", "asset type USD:citibank")
	expectOutput(t, cli, "credit_alphanum12", "asset type USD:citibank:credit_alphanum12")
}
