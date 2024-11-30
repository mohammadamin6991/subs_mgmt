
CREATE TABLE public.plans (
    id INTEGER DEFAULT nextval('public.plan_id_seq'::regclass) NOT NULL,
    name VARCHAR(255),
    description VARCHAR(255),
    days_per_interval INTEGER,
    storage_type VARCHAR(255),
    storage_size INTEGER,
    price INTEGER,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    UNIQUE (name)
);


ALTER TABLE public.plans OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.plan_id_seq', 1, true);


-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.plans
    ADD CONSTRAINT plans_pkey PRIMARY KEY (id);
