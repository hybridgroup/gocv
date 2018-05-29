package ie

import (
	"os"
	"testing"
)

func TestInferenceEnginePlugin(t *testing.T) {
	pd := NewInferenceEnginePluginDispatcher(DefaultLibPath())
	defer pd.Close()

	pu := pd.GetPluginByDevice("CPU")
	defer pu.Close()
}

func TestCNNNetwork(t *testing.T) {
	nr := NewCNNNetReader()
	defer nr.Close()

	modelPath := os.Getenv("INTEL_CVSDK_DIR") + "/deployment_tools/intel_models/"
	binPath := modelPath + "age-gender-recognition-retail-0013/FP32/age-gender-recognition-retail-0013.xml"
	weightsPath := modelPath + "age-gender-recognition-retail-0013/FP32/age-gender-recognition-retail-0013.bin"

	nr.ReadNetwork(binPath)
	nr.ReadWeights(weightsPath)

	net := nr.GetNetwork()
	defer net.Close()

	if net.Size() != 25 {
		t.Errorf("Invalid CNNNet: %d", net.Size())
	}
}
