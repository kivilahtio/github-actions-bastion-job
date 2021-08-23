Testing GitHub Actions here to do some job dispatching based on PR contents.

This is to protect self-hosted Runners
    https://docs.github.com/en/actions/hosting-your-own-runners/using-self-hosted-runners-in-a-workflow

from fork-and-PR attacks trying to access the self-hosted persistent build environment.

As the routing is based on non-easily-spoofable event payload attributes, this should be reasonably safe.
    https://docs.github.com/en/actions/learn-github-actions/security-hardening-for-github-actions#understanding-the-risk-of-script-injections
