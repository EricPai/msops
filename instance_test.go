package msops

import "testing"

func TestRegisterAndUnRegister(t *testing.T) {
	if err := Register(testEndpoint1, testDBAUser, testDBAPass, testReplUser, testReplPass, testParams); err != nil {
		t.Errorf("Register endpoint1 error: %s", err.Error())
	}
	if err := Register(testEndpoint2, testDBAUser, testDBAPass, testReplUser, testReplPass, testParams); err != nil {
		t.Errorf("Register endpoint2 error: %s", err.Error())
	}
	if err := Register(testEndpoint3, testDBAUser, testDBAPass, testReplUser, testReplPass, testParams); err != nil {
		t.Errorf("Register endpoint3 error: %s", err.Error())
	}
	if err := Register(badEndpoint, testDBAUser, testDBAPass, testReplUser, testReplPass, testParams); err != nil {
		t.Errorf("Register badEndpoint error: %s", err.Error())
	}

	Unregister(testEndpoint1)
	if len(connectionPool) != 3 || connectionPool[testEndpoint1] != nil {
		t.Errorf("Unregister endpoint1 error")
	}
	if err := Register(testEndpoint1, testDBAUser, testDBAPass, testReplUser, testReplPass, testParams); err != nil {
		t.Errorf("Reregister endpoint1 error: %s", err.Error())
	}

}

func TestCheckInstanceStatus(t *testing.T) {
	if inst := CheckInstance(testEndpoint1); inst != InstanceOK {
		t.Errorf("Endpoint1 instance status error: actual %d, expect %d", inst, InstanceOK)
	}
	if inst := CheckInstance(badEndpoint); inst != InstanceERROR {
		t.Errorf("BadEndpoint instance status error: actual %d, expect %d", inst, InstanceERROR)
	}
	if inst := CheckInstance(unregisteredEndpoint); inst != InstanceUnregistered {
		t.Errorf("UnregisteredEndpoint instance status error: actual %d, expect %d", inst, InstanceUnregistered)
	}
}
