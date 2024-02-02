package geospatials

import (
	"encoding/json"
	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/rofinafiin/gcfbackend"
	"net/http"
	"os"
)

func GeoIntersectcloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(RequestGeoIntersects)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
		if !existing {
			resp.Status = false
			resp.Message = "Kamu kayaknya belum punya akun"
		} else {
			datageo, err := GeoIntersectQuery(conn, req.Coordinates)
			if err != nil {
				resp.Status = false
				resp.Message = err.Error()
			}
			jsoncihuy, _ := json.Marshal(datageo)
			resp.Status = true
			resp.Message = "Data Berhasil diambil"
			resp.Token = string(jsoncihuy)
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}

func GeoNearcloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(RequestGeoIntersects)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
		if !existing {
			resp.Status = false
			resp.Message = "Kamu kayaknya belum punya akun"
		} else {
			datageo, err := GeoNearQuery(conn, req.Coordinates, 100)
			if err != nil {
				resp.Status = false
				resp.Message = err.Error()
			}
			jsoncihuy, _ := json.Marshal(datageo)
			resp.Status = true
			resp.Message = "Data Berhasil diambil"
			resp.Token = string(jsoncihuy)
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}

func GeoWithincloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(RequestGeoIntersects)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Message = "error parsing application/json: " + err.Error()
		} else {
			existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
			if !existing {
				resp.Status = false
				resp.Message = "Kamu kayaknya belum punya akun"
			} else {
				datageo, err := GeoWithinQuery(conn, req.Coordinates)
				if err != nil {
					resp.Status = false
					resp.Message = err.Error()
				}
				jsoncihuy, _ := json.Marshal(datageo)
				resp.Status = true
				resp.Message = "Data Berhasil diambil"
				resp.Token = string(jsoncihuy)
			}
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}

func GeoNearspherecloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(Nearspherereq)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Message = "error parsing application/json: " + err.Error()
		} else {
			existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
			if !existing {
				resp.Status = false
				resp.Message = "Kamu kayaknya belum punya akun"
			} else {
				datageo, err := GeoNearSphereQuery(conn, req.Coordinates, req.Radius)
				if err != nil {
					resp.Status = false
					resp.Message = err.Error()
				}
				jsoncihuy, _ := json.Marshal(datageo)
				resp.Status = true
				resp.Message = "Data Berhasil diambil"
				resp.Token = string(jsoncihuy)
			}
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}

func GeoBoxcloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(GeoBoxReq)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Message = "error parsing application/json: " + err.Error()
		} else {
			existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
			if !existing {
				resp.Status = false
				resp.Message = "Kamu kayaknya belum punya akun"
			} else {
				datageo, err := GeoBoxQuery(conn, req.LowerLeft, req.LowerRight)
				if err != nil {
					resp.Status = false
					resp.Message = err.Error()
				}
				jsoncihuy, _ := json.Marshal(datageo)
				resp.Status = true
				resp.Message = "Data Berhasil diambil"
				resp.Token = string(jsoncihuy)
			}
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}

func GeoCentercloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(Nearspherereq)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Message = "error parsing application/json: " + err.Error()
		} else {
			existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
			if !existing {
				resp.Status = false
				resp.Message = "Kamu kayaknya belum punya akun"
			} else {
				datageo, err := GeoCenterQuery(conn, req.Coordinates, req.Radius)
				if err != nil {
					resp.Status = false
					resp.Message = err.Error()
				}
				jsoncihuy, _ := json.Marshal(datageo)
				resp.Status = true
				resp.Message = "Data Berhasil diambil"
				resp.Token = string(jsoncihuy)
			}
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}

func GeoGeometrycloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(Geometryreq)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Message = "error parsing application/json: " + err.Error()
		} else {
			existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
			if !existing {
				resp.Status = false
				resp.Message = "Kamu kayaknya belum punya akun"
			} else {
				datageo, err := GeoGeometryQuery(conn, req.Geometry)
				if err != nil {
					resp.Status = false
					resp.Message = err.Error()
				}
				jsoncihuy, _ := json.Marshal(datageo)
				resp.Status = true
				resp.Message = "Data Berhasil diambil"
				resp.Token = string(jsoncihuy)
			}
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}

func GeoMaxDistancecloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(Nearspherereq)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Message = "error parsing application/json: " + err.Error()
		} else {
			existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
			if !existing {
				resp.Status = false
				resp.Message = "Kamu kayaknya belum punya akun"
			} else {
				datageo, err := GeoMaxDistanceQuery(conn, req.Coordinates, req.Radius)
				if err != nil {
					resp.Status = false
					resp.Message = err.Error()
				}
				jsoncihuy, _ := json.Marshal(datageo)
				resp.Status = true
				resp.Message = "Data Berhasil diambil"
				resp.Token = string(jsoncihuy)
			}
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}

func GeoMinDistancecloud(publickey, Mongostring, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	req := new(Nearspherereq)
	conn := pasproj.MongoCreateConnection(Mongostring, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		resp.Status = false
		resp.Message = "Header Login Not Exist"
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Message = "error parsing application/json: " + err.Error()
		} else {
			existing := gcfbackend.IsExist(tokenlogin, os.Getenv(publickey))
			if !existing {
				resp.Status = false
				resp.Message = "Kamu kayaknya belum punya akun"
			} else {
				datageo, err := GeoMinDistanceQuery(conn, req.Coordinates, req.Radius)
				if err != nil {
					resp.Status = false
					resp.Message = err.Error()
				}
				jsoncihuy, _ := json.Marshal(datageo)
				resp.Status = true
				resp.Message = "Data Berhasil diambil"
				resp.Token = string(jsoncihuy)
			}
		}
	}
	return gcfbackend.ReturnStringStruct(resp)
}
