-- name: CreateDispatcher :execresult
INSERT INTO
    dispatchers (id, request_id, registered_number_shipper, registered_number_dispatcher, zipcode_origin, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetDispatcherByID :one
SELECT
    *
FROM dispatchers
WHERE id = ?;

-- name: CreateOffer :execresult
INSERT INTO
    offers (id, dispatcher_id, offer, simulation_type, carrier_id, service, service_code, service_description, delivery_time, original_delivery_time, identifier, delivery_note, home_delivery, carrier_needs_to_return_to_sender, expiration, cost_price, final_price, weights, composition, esg, modal, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetOfferByID :one
SELECT
    *
FROM offers
WHERE id = ?;

-- name: CreateCarrier :execresult
INSERT INTO
    carriers (reference, name, registered_number, state_inscription, logo_url, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetCarrierByID :one
SELECT
    *
FROM carriers
WHERE reference = ?;