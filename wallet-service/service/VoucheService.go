package service

import (
	"grabathon/api/request_response"
	"grabathon/models"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

func CreateVouches(vouches []request_response.Vouch) []request_response.Vouch {
	var vchs []models.Vouch
	for _,v := range vouches{
		vchs = append(vchs, request_response.GetVouchModel(v))
	}
	vchs, err := models.CreateVouches(vchs)
	if err != nil {
		return []request_response.Vouch{}
	}

	var vchsResponse []request_response.Vouch
	for _,v := range vchs{
		vchsResponse = append(vchsResponse, request_response.GetVouchRequestResponse(v))
	}
	return vchsResponse
}

func CreateVouch(vouch request_response.Vouch) request_response.Vouch {
	vch, err := models.CreateVouch(request_response.GetVouchModel(vouch))
	if err != nil {
		return request_response.Vouch{}
	}

	return request_response.GetVouchRequestResponse(vch)
}

func GetVouch(id int) request_response.Vouch {
	vch, err := models.GetVouch(id)
	if err != nil {
		return request_response.Vouch{}
	}
	return request_response.GetVouchRequestResponse(vch)
}

func GetAllVouchesForVouchee(voucheeId string) []request_response.Vouch {
	vches, err := models.GetAllVouchesForVouchee(voucheeId)
	if err != nil {
		return []request_response.Vouch{}
	}

	var vouchesResponse []request_response.Vouch
	for _,v := range vches {
		vouchesResponse = append(vouchesResponse, request_response.GetVouchRequestResponse(v))
	}

	return vouchesResponse
}

func GetAllVouchesForVoucher(voucherId string) []request_response.Vouch {
	vches, err := models.GetAllVouchesForVoucher(voucherId)
	if err != nil {
		return []request_response.Vouch{}
	}

	var vouchesResponse []request_response.Vouch
	for _,v := range vches {
		vouchesResponse = append(vouchesResponse, request_response.GetVouchRequestResponse(v))
	}

	return vouchesResponse
}

func DeleteVouch(id int) error {
	v := models.Vouch{
		Id:        id,
		VoucheeId: "",
		VoucherId: "",
	}
	err := models.DeleteVouch(v)
	return err
}

func UpdateVouch(id int, status string) (request_response.Vouch,error) {
	vch := GetVouch(id)
	vch.Status = status
	v, err := models.UpdateVouch(request_response.GetVouchModel(vch))
	return request_response.GetVouchRequestResponse(v),err
}