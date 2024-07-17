-- name: CreateDispatcher :execresult
INSERT INTO
    dispatchers (id, request_id, registered_number_shipper, registered_number_dispatcher, zipcode_origin, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetDispatcherByID :one
SELECT
    *
FROM dispatchers
WHERE id = ?;

-- name: GetDispatcherIdsWithLimit :many
SELECT
    id
FROM dispatchers
ORDER BY created_at DESC
LIMIT ?;

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

-- name: GetPriceMetric :one
SELECT
    CAST(MIN(final_price) AS DECIMAL (13,2)) AS cheaper_shipping,
    CAST(MAX(final_price) AS DECIMAL (13,2)) AS most_expensive_shipping
FROM offers
ORDER BY created_at;

-- name: GetCarrierMetric :many
SELECT
    carriers.name AS carrier_name,
    COUNT(carrier_id) AS total,
    CAST(SUM(final_price) AS DECIMAL (13,2))  AS final_price_sum,
    CAST(AVG(final_price) AS DECIMAL (13,2)) AS final_price_mean
FROM offers
    JOIN carriers ON offers.carrier_id = carriers.reference
GROUP BY carrier_id ORDER BY carrier_name;

-- name: GetPriceMetricWithLimit :one
SELECT
    CAST(MIN(r.final_price) AS DECIMAL (13,2)) AS cheaper_shipping,
    CAST(MAX(r.final_price) AS DECIMAL (13,2)) AS most_expensive_shipping
FROM (
     SELECT
         offers.final_price
     FROM offers
     WHERE offers.dispatcher_id IN (sqlc.slice('ids'))
     ORDER BY offers.created_at
 ) AS r;

-- name: GetCarrierMetricsWithLimit :many
SELECT
    r.carrier_name,
    COUNT(r.carrier_id) AS total,
    CAST(SUM(r.final_price) AS DECIMAL (13,2))  AS final_price_sum,
    CAST(AVG(r.final_price) AS DECIMAL (13,2)) AS final_price_mean
FROM (
     SELECT
         offers.carrier_id,
         carriers.name AS carrier_name,
         offers.final_price
     FROM offers
              JOIN carriers ON offers.carrier_id = carriers.reference
     WHERE offers.dispatcher_id IN (sqlc.slice('ids'))
     ORDER BY offers.created_at
 ) AS r GROUP BY r.carrier_name ORDER BY r.carrier_name;


