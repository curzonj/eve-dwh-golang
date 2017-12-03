package model

import (
	"github.com/pkg/errors"
)

func (d *DAO) GetPlanetarySchematicOutputID(schematicID int32) (int32, error) {
	var outputID int32
	err := d.QueryRow("select output from sde_planetary_schematics where schematic_id = $1", schematicID).Scan(&outputID)
	if err != nil {
		return outputID, errors.Wrap(err, "GetPlanetarySchematicOutputID")
	}

	return outputID, nil
}

func (d *DAO) GetPlanetarySchematicInputQuantity(schematicID int32, typeID int32) (int64, error) {
	var schematicQuantity int64
	err := d.QueryRow("select contents->'inputs'->$1 from sde_planetary_schematics where schematic_id = $2", typeID, schematicID).Scan(&schematicQuantity)
	if err != nil {
		return schematicQuantity, errors.Wrap(err, "GetPlanetarySchematicInputQuantity")
	}

	return schematicQuantity, nil
}

type SDEType struct {
	Name   string  `db:"name"`
	Volume float32 `db:"volume"`
}

func (d *DAO) GetSDEType(typeID int32) (*SDEType, error) {
	var t SDEType
	err := d.Get(&t, "select name, volume from sde_types where type_id = $1 limit 1", typeID)
	if err != nil {
		return nil, errors.Wrap(err, "GetSDEType")
	}

	return &t, nil
}
