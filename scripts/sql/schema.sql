DROP TABLE IF EXISTS offers;
DROP TABLE IF EXISTS carriers;
DROP TABLE IF EXISTS dispatchers;

CREATE TABLE dispatchers
(
    id                           VARCHAR(255) PRIMARY KEY,
    request_id                   VARCHAR(255) NOT NULL,
    registered_number_shipper    VARCHAR(255) NOT NULL,
    registered_number_dispatcher VARCHAR(255) NOT NULL,
    zipcode_origin               INT          NOT NULL,
    created_at                   TIMESTAMP    NOT NULL,
    updated_at                   TIMESTAMP    NOT NULL
);

CREATE TABLE carriers
(
    reference         BIGINT PRIMARY KEY,
    name              VARCHAR(255) NOT NULL,
    registered_number VARCHAR(255) NOT NULL,
    state_inscription VARCHAR(255) NOT NULL,
    logo_url          VARCHAR(255) NOT NULL,
    created_at        TIMESTAMP    NOT NULL,
    updated_at        TIMESTAMP    NOT NULL
);

CREATE TABLE offers
(
    id                                VARCHAR(255) PRIMARY KEY,
    dispatcher_id                     VARCHAR(255) NOT NULL,
    offer                             INT          NOT NULL,
    simulation_type                   INT          NOT NULL,
    carrier_id                        BIGINT       NOT NULL,
    service                           VARCHAR(255) NOT NULL,
    service_code                      VARCHAR(255),
    service_description               VARCHAR(255),
    delivery_time                     JSON         NOT NULL,
    original_delivery_time            JSON         NOT NULL,
    identifier                        VARCHAR(255),
    delivery_note                     VARCHAR(255),
    home_delivery                     BOOLEAN,
    carrier_needs_to_return_to_sender BOOLEAN,
    expiration                        DATETIME     NOT NULL,
    cost_price                        DECIMAL      NOT NULL,
    final_price                       DECIMAL      NOT NULL,
    weights                           JSON         NOT NULL,
    composition                       JSON,
    esg                               JSON,
    modal                             VARCHAR(255),
    created_at                        TIMESTAMP    NOT NULL,
    updated_at                        TIMESTAMP    NOT NULL,

    FOREIGN KEY (carrier_id) REFERENCES carriers (reference),
    FOREIGN KEY (dispatcher_id) REFERENCES dispatchers (id)
);

