package service

import (
	"bytes"
	"e-backend-boilerplate/modules/cv/models"
	"e-backend-boilerplate/modules/cv/repository"
	"path"
	"text/template"
)

type Service struct {
	repo            *repository.Repository
	cvBaseURL       string
	cvTemplatesPath string
}

type CVData struct {
	Education    []models.CVEducationItem
	Experience   []models.CVExperienceItem
	Publications []models.CVPublication
}

func NewService(repo *repository.Repository, cvBaseURL, cvTemplatesPath string) *Service {
	return &Service{repo, cvBaseURL, cvTemplatesPath}
}

func (s *Service) GetDevTimeline(reverse bool) ([]models.DevTimelineItem, error) {
	return s.repo.GetDevTimeline(reverse)
}

func (s *Service) GetCVCommon() (*models.CVCommon, error) {
	return s.repo.GetCVCommon()
}

func (s *Service) GetCVEducation() ([]models.CVEducationItem, error) {
	return s.repo.GetCVEducation()
}

func (s *Service) GetCVExperience() ([]models.CVExperienceItem, error) {
	return s.repo.GetCVExperience()
}

func (s *Service) GetCVPublications() ([]models.CVPublication, error) {
	list, err := s.repo.GetCVPublications()
	if err != nil {
		return nil, err
	}

	for key, item := range list {
		item.ImageURL = s.cvBaseURL + item.ImageURL
		list[key] = item
	}

	return list, nil
}

func (s *Service) GetCVSoftwareProjects() ([]models.CVSoftwareProject, error) {
	list, err := s.repo.GetCVSoftwareProjects()

	if err != nil {
		return nil, err
	}

	for key, item := range list {
		item.ImageURL = s.cvBaseURL + item.ImageURL
		list[key] = item
	}

	return list, nil
}

func (s *Service) GetCVLatex() (string, error) {
	// Retrieve all necessary CV data
	education, err := s.GetCVEducation()
	if err != nil {
		return "", err
	}

	experience, err := s.GetCVExperience()
	if err != nil {
		return "", err
	}

	publications, err := s.GetCVPublications()
	if err != nil {
		return "", err
	}

	cvData := CVData{
		Education:    education,
		Experience:   experience,
		Publications: publications,
	}

	// Prepare LaTeX template
	tmplPath := path.Join(s.cvTemplatesPath, "cv.tex.tmpl")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, cvData)
	if err != nil {
		return "", err
	}

	renderedCV := buf.String()

	return renderedCV, nil
}
