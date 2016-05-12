package api

import (
	"encoding/json"
	"errors"
	gostErrors "github.com/geodan/gost/src/errors"
	"github.com/geodan/gost/src/sensorthings/entities"
	"github.com/geodan/gost/src/sensorthings/models"
	"github.com/geodan/gost/src/sensorthings/odata"
)

// GetObservation returns an observation by id
func (a *APIv1) GetObservation(id string, qo *odata.QueryOptions) (*entities.Observation, error) {
	o, err := a.db.GetObservation(id)
	if err != nil {
		return nil, err
	}

	o.SetLinks(a.config.GetExternalServerURI())
	return o, nil
}

// GetObservations return all observations
func (a *APIv1) GetObservations(qo *odata.QueryOptions) (*models.ArrayResponse, error) {
	observations, err := a.db.GetObservations()
	return processObservations(a, observations, err)
}

// GetObservationsByDatastream todo
func (a *APIv1) GetObservationsByDatastream(datastreamID string, qo *odata.QueryOptions) (*models.ArrayResponse, error) {
	return nil, gostErrors.NewRequestNotImplemented(errors.New("not implemented yet"))
}

func processObservations(a *APIv1, datastreams []*entities.Observation, err error) (*models.ArrayResponse, error) {
	if err != nil {
		return nil, err
	}

	uri := a.config.GetExternalServerURI()
	for idx, item := range datastreams {
		i := *item
		i.SetLinks(uri)
		datastreams[idx] = &i
	}

	var data interface{} = datastreams
	return &models.ArrayResponse{
		Count: len(datastreams),
		Data:  &data,
	}, nil
}

// PostObservation todo
func (a *APIv1) PostObservation(observation entities.Observation) (*entities.Observation, []error) {
	_, err := observation.ContainsMandatoryParams()
	if err != nil {
		return nil, err
	}

	//ToDo check for linked featureofinterest
	no, err2 := a.db.PostObservation(observation)
	if err2 != nil {
		return nil, []error{err2}
	}

	no.SetLinks(a.config.GetExternalServerURI())

	json, _ := json.Marshal(no)
	s := string(json)

	//ToDo: TEST
	a.mqtt.Publish("Datastreams(1)/Observations", s, 0)
	a.mqtt.Publish("Observations", s, 0)

	return no, nil
}

// PostObservationByDatastream creates a Datastream with given id for the Observation and calls PostObservation
func (a *APIv1) PostObservationByDatastream(datastreamID string, observation entities.Observation) (*entities.Observation, []error) {
	observation.Datastream = &entities.Datastream{ID: datastreamID}
	return a.PostObservation(observation)
}

// PatchObservation todo
func (a *APIv1) PatchObservation(id string, observation entities.Observation) (*entities.Observation, error) {
	return nil, gostErrors.NewRequestNotImplemented(errors.New("not implemented yet"))
}

// DeleteObservation todo
func (a *APIv1) DeleteObservation(id string) error {
	return gostErrors.NewRequestNotImplemented(errors.New("not implemented yet"))
}