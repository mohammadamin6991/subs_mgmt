
CREATE TABLE public.instances (
    id INTEGER DEFAULT nextval('public.instance_id_seq'::regclass) NOT NULL,
    plan_id VARCHAR(255),
    user_id VARCHAR(255),
    is_active INTEGER,
    access_key VARCHAR(255),
    secret_key VARCHAR(255),
    endpoint VARCHAR(255),
    region VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE
);


ALTER TABLE public.instances OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.instance_id_seq', 1, true);


-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.instances
    ADD CONSTRAINT instances_pkey PRIMARY KEY (id);
