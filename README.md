## AssertDB

Assert DB is a Golang library to facilitate the work around database assertions when developing integraitons tests.

This lib builds on top of the `testfy/suite` so that all the setup that needs to happen before the tests, stays in the `SetupSuite` and the `TearDownSuite` methods
