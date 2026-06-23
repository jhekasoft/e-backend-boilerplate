package repository

import (
	"e-backend-boilerplate/modules/cv/models"
	"os"
	"path"
	"slices"

	"gopkg.in/yaml.v3"
)

type Repository struct {
	cvDataPath string
}

func NewRepository(cvDataPath string) *Repository {
	return &Repository{cvDataPath}
}

func (r *Repository) GetDevTimeline(reverse bool) (list []models.DevTimelineItem, err error) {
	yamlFile, err := os.ReadFile(path.Join(r.cvDataPath, "developer-timeline/list.yml"))
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &list)
	if err != nil {
		return
	}

	if reverse {
		slices.Reverse(list)
	}

	return
}

func (r *Repository) GetCVCommon() (data *models.CVCommon, err error) {
	yamlFile, err := os.ReadFile(path.Join(r.cvDataPath, "cv/common.yml"))
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return
	}

	return
}

func (r *Repository) GetCVEducation() (list []models.CVEducationItem, err error) {
	yamlFile, err := os.ReadFile(path.Join(r.cvDataPath, "cv/education.yml"))
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &list)
	if err != nil {
		return
	}
	slices.Reverse(list)

	return
}

func (r *Repository) GetCVExperience() (list []models.CVExperienceItem, err error) {
	yamlFile, err := os.ReadFile(path.Join(r.cvDataPath, "cv/experience.yml"))
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &list)
	if err != nil {
		return
	}
	slices.Reverse(list)

	return
}

func (r *Repository) GetCVPublications() (list []models.CVPublication, err error) {
	yamlFile, err := os.ReadFile(path.Join(r.cvDataPath, "cv/publications.yml"))
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &list)
	if err != nil {
		return
	}
	slices.Reverse(list)

	return
}

func (r *Repository) GetCVSoftwareProjects() (list []models.CVSoftwareProject, err error) {
	yamlFile, err := os.ReadFile(path.Join(r.cvDataPath, "cv/software-projects.yml"))
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &list)
	if err != nil {
		return
	}
	slices.Reverse(list)

	return
}
