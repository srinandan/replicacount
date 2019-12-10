package replicacount

import (
	"os"

	types "github.com/srinandan/replicacount/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetReplicaCount(runtimeReplicaSetName string) (types.CountResponse, error) {

	countResponse := types.CountResponse{}
	namespace := os.Getenv("APIGEE_NAMESPACE")

	config, err := rest.InClusterConfig()
	if err != nil {
		return countResponse, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return countResponse, err
	}

	replicaSetList, err := clientset.AppsV1().ReplicaSets(namespace).List(metav1.ListOptions{})
	if err != nil {
		return countResponse, err
	}

	for _, replicaSet := range replicaSetList.Items {
		if runtimeReplicaSetName == replicaSet.ObjectMeta.Name {
			countResponse.Replicas = replicaSet.Status.Replicas
			countResponse.ReadyReplicas = replicaSet.Status.ReadyReplicas
			countResponse.AvailableReplicas = replicaSet.Status.AvailableReplicas
			return countResponse, nil
		}
	}

	return countResponse, nil
}
