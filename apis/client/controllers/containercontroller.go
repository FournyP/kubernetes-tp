package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ContainerController struct {
	kubernetesClient *kubernetes.Clientset
}

func NewContainerController() *ContainerController {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the kubernetesClient
	kubernetesClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &ContainerController{
		kubernetesClient: kubernetesClient,
	}
}

func (controller *ContainerController) CreateContainer(c *fiber.Ctx) error {
	name := c.Query("name", "")

	if name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Please provide a name for the service")
	}

	createDeployment(name, controller.kubernetesClient)

	return c.SendString("Container created !")
}

func renderTemplate(templatePath string, variables map[string]string) []byte {
	data, err := os.ReadFile(templatePath)

	if err != nil {
		panic(err)
	}

	str := string(data)

	for key, element := range variables {
		str = strings.Replace(str, "{{ "+key+" }}", element, -1)
	}

	return []byte(str)
}

func createDeployment(name string, kubernetesClient *kubernetes.Clientset) {
	deploymentsClient := kubernetesClient.AppsV1().Deployments(corev1.NamespaceDefault)

	variables := make(map[string]string)

	variables["NAME"] = "test"
	data := renderTemplate("./deployment.template.json", variables)

	var deployment appsv1.Deployment

	err := json.Unmarshal(data, &deployment)

	if err != nil {
		panic(err)
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), &deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}
