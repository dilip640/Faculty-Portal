package storage

import log "github.com/sirupsen/logrus"

// InsertRoute will add route in table
func InsertRoute(route Route) error {
	sqlStatement := `INSERT INTO application_route (applier, route_from, route_to, ccf_post) 
						VALUES ($1, $2, $3, $4);`
	_, err := db.Exec(sqlStatement, route.Applier, route.RouteFrom, route.RouteTo, route.CCFPost)
	return err
}

// DeleteRoute remove route
func DeleteRoute(routeID int) error {
	sqlStatement := `DELETE FROM application_route WHERE id = $1;`
	_, err := db.Exec(sqlStatement, routeID)
	return err
}

// GetRouteStatus returns all the routes corresponding to that person
func GetRouteStatus(routeTo string) ([]*Route, error) {
	routes := make([]*Route, 0)

	rows, err := db.Query(
		`SELECT id, applier, route_from, route_to, ccf_post 
			FROM application_route WHERE route_to = $1`, routeTo)
	if err != nil {
		return routes, err
	}
	defer rows.Close()

	for rows.Next() {
		var route Route

		if err := rows.Scan(&route.ID, &route.Applier, &route.RouteFrom, &route.RouteTo,
			&route.CCFPost); err == nil {
			routes = append(routes, &route)
		} else {
			log.Error(err)
		}
	}

	return routes, nil
}

// GetAllRoutes returns all routes
func GetAllRoutes() ([]*Route, error) {
	routes := make([]*Route, 0)

	rows, err := db.Query(
		`SELECT id, applier, route_from, route_to, ccf_post FROM application_route`)
	if err != nil {
		return routes, err
	}
	defer rows.Close()

	for rows.Next() {
		var route Route

		if err := rows.Scan(&route.ID, &route.Applier, &route.RouteFrom, &route.RouteTo,
			&route.CCFPost); err == nil {
			routes = append(routes, &route)
		} else {
			log.Error(err)
		}
	}

	return routes, nil
}

// Route struct
type Route struct {
	ID        int    `json:"id"`
	Applier   string `json:"applier"`
	RouteTo   string `json:"route_to"`
	RouteFrom string `json:"route_from"`
	CCFPost   string `json:"ccf_post"`
}
