--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.2 (Debian 17.2-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: accessoire_defaut; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accessoire_defaut (
    id integer NOT NULL,
    categorie_id integer,
    accessoire_id integer
);


ALTER TABLE public.accessoire_defaut OWNER TO postgres;

--
-- Name: accessoire_defaut_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.accessoire_defaut_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.accessoire_defaut_id_seq OWNER TO postgres;

--
-- Name: accessoire_defaut_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.accessoire_defaut_id_seq OWNED BY public.accessoire_defaut.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    nom character varying(255) NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: compatibilites; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.compatibilites (
    id integer NOT NULL,
    categorie_id integer,
    accessoire_id integer
);


ALTER TABLE public.compatibilites OWNER TO postgres;

--
-- Name: compatibilites_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.compatibilites_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.compatibilites_id_seq OWNER TO postgres;

--
-- Name: compatibilites_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.compatibilites_id_seq OWNED BY public.compatibilites.id;


--
-- Name: liste_accessoire; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.liste_accessoire (
    id integer NOT NULL,
    nom character varying(255) NOT NULL,
    image character varying(255),
    prix numeric(10,2) DEFAULT 10 NOT NULL
);


ALTER TABLE public.liste_accessoire OWNER TO postgres;

--
-- Name: liste_accessoire_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.liste_accessoire_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.liste_accessoire_id_seq OWNER TO postgres;

--
-- Name: liste_accessoire_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.liste_accessoire_id_seq OWNED BY public.liste_accessoire.id;


--
-- Name: moyens_roulants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.moyens_roulants (
    id integer NOT NULL,
    nom character varying(255) NOT NULL,
    categorie_id integer,
    roues integer NOT NULL,
    emplacement character varying(255),
    type_base character varying(50),
    taille character varying(5),
    departement character varying(255),
    image character varying(255),
    prix numeric(10,2) NOT NULL
);


ALTER TABLE public.moyens_roulants OWNER TO postgres;

--
-- Name: moyens_roulants_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.moyens_roulants_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.moyens_roulants_id_seq OWNER TO postgres;

--
-- Name: moyens_roulants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.moyens_roulants_id_seq OWNED BY public.moyens_roulants.id;


--
-- Name: accessoire_defaut id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accessoire_defaut ALTER COLUMN id SET DEFAULT nextval('public.accessoire_defaut_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: compatibilites id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.compatibilites ALTER COLUMN id SET DEFAULT nextval('public.compatibilites_id_seq'::regclass);


--
-- Name: liste_accessoire id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.liste_accessoire ALTER COLUMN id SET DEFAULT nextval('public.liste_accessoire_id_seq'::regclass);


--
-- Name: moyens_roulants id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.moyens_roulants ALTER COLUMN id SET DEFAULT nextval('public.moyens_roulants_id_seq'::regclass);


--
-- Data for Name: accessoire_defaut; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.accessoire_defaut (id, categorie_id, accessoire_id) FROM stdin;
1	1	13
2	1	15
3	2	2
4	2	13
5	2	15
6	3	2
7	3	13
8	3	15
9	4	10
10	4	11
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, nom) FROM stdin;
1	Plateau roulant
2	Chariot
3	Base roulante
4	Chariot AGV
\.


--
-- Data for Name: compatibilites; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.compatibilites (id, categorie_id, accessoire_id) FROM stdin;
1	1	1
2	1	3
3	1	4
4	1	5
5	1	6
6	1	8
7	1	9
8	1	11
9	1	14
10	2	5
11	2	6
12	2	7
13	2	8
14	2	9
15	2	11
16	2	14
17	3	3
18	3	4
19	3	5
20	3	6
21	3	7
22	3	8
23	3	9
24	3	11
25	3	12
26	3	14
27	4	2
28	4	5
29	4	6
30	4	7
31	4	8
32	4	9
\.


--
-- Data for Name: liste_accessoire; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.liste_accessoire (id, nom, image, prix) FROM stdin;
3	Attelage à main	Attelageamain.png	10.00
5	Avaloirs	avaloirs.png	10.00
6	Rallonge AV AR	RallongeAVAR.png	10.00
7	Support étiquette barre	supportetiquettebarre.png	10.00
8	Support étiquette châssis	supportetiquettechassis.png	10.00
10	Interface AGV	interfaceAGV.png	10.00
11	Rehausse AV centre AR	rehausses.png	10.00
14	Patin immobilisation	patinimmobilisation.png	10.00
13	Anti basculeur	antibasculeur.png	10.00
4	Poignet flexible	poignetflexible.png	10.00
1	Timon	Timon.png	10.00
9	Support avaloir central	supportavaloircentral.png	10.00
12	Pion easy mover	pioneasymover.png	10.00
15	Plateau Rehausse	plateaurehausse.png	10.00
2	Barre de manœuvre	Barredemanoeuvre.png	10.00
\.


--
-- Data for Name: moyens_roulants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.moyens_roulants (id, nom, categorie_id, roues, emplacement, type_base, taille, departement, image, prix) FROM stdin;
30	Chariot AGV frontal uni directionnel osmoze	4	4	Coin	Plein	M	Tôlerie	ChariotAGV2.png	100.00
2	Plateau roulant	1	4	Coin	Plein	M	Tôlerie	Plateauroulant.png	100.00
31	Chariot AGV frontal uni directionnel osmoze	4	4	Coin	Plein	L	Tôlerie	ChariotAGV2.png	100.00
4	Chariot 4 roues au coin	2	4	Coin	Plein	S	Tôlerie	Chariot1.png	100.00
5	Chariot 4 roues au coin	2	4	Coin	Plein	M	Tôlerie	Chariot1.png	100.00
6	Chariot 4 roues au coin	2	4	Coin	Plein	L	Tôlerie	Chariot1.png	100.00
7	Chariot 6 roues	2	6	Coin	Plein	S	Tôlerie	Chariot4.png	100.00
8	Chariot 6 roues	2	6	Coin	Plein	M	Tôlerie	Chariot4.png	100.00
9	Chariot 6 roues	2	6	Coin	Plein	L	Tôlerie	Chariot4.png	100.00
10	Chariot 4 roues au coin	2	4	Coin	Échancré	M	Tôlerie	Chariot2.png	100.00
11	Chariot 4 roues au coin	2	4	Coin	Échancré	L	Tôlerie	Chariot2.png	100.00
12	Chariot 4 roues losange	2	4	Losange	Plein	S	Tôlerie	Chariot3.png	100.00
13	Chariot 4 roues losange	2	4	Losange	Plein	M	Tôlerie	Chariot3.png	100.00
14	Chariot 4 roues losange	2	4	Losange	Plein	L	Tôlerie	Chariot3.png	100.00
15	Base roulante 4 roues au coin	3	4	Coin	Plein	S	Tôlerie	Baseroulante1.png	100.00
16	Base roulante 4 roues au coin	3	4	Coin	Plein	M	Tôlerie	Baseroulante1.png	100.00
17	Base roulante 4 roues au coin	3	4	Coin	Plein	L	Tôlerie	Baseroulante1.png	100.00
18	Base roulante 6 roues	3	6	Coin	Plein	S	Tôlerie	Baseroulante4.png	100.00
19	Base roulante 6 roues	3	6	Coin	Plein	M	Tôlerie	Baseroulante4.png	100.00
20	Base roulante 6 roues	3	6	Coin	Plein	L	Tôlerie	Baseroulante4.png	100.00
21	Base roulante 4 roues au coin	3	4	Coin	Échancré	M	Tôlerie	Baseroulante2.png	100.00
22	Base roulante 4 roues au coin	3	4	Coin	Échancré	L	Tôlerie	Baseroulante2.png	100.00
23	Base roulante 4 roues losange	3	4	Losange	Plein	S	Tôlerie	Baseroulante3.png	100.00
24	Base roulante 4 roues losange	3	4	Losange	Plein	M	Tôlerie	Baseroulante3.png	100.00
25	Base roulante 4 roues losange	3	4	Losange	Plein	L	Tôlerie	Baseroulante3.png	100.00
26	Chariot AGV central osmoze	4	4	Coin	Plein	S	Tôlerie	ChariotAGV1.png	100.00
27	Chariot AGV central osmoze	4	4	Coin	Plein	M	Tôlerie	ChariotAGV1.png	100.00
28	Chariot AGV central osmoze	4	4	Coin	Plein	L	Tôlerie	ChariotAGV1.png	100.00
29	Chariot AGV frontal uni directionnel osmoze	4	4	Coin	Plein	S	Tôlerie	ChariotAGV2.png	100.00
32	Chariot AGV frontal bi directionnel osmoze	4	4	Coin	Plein	S	Tôlerie	ChariotAGV3.png	100.00
33	Chariot AGV frontal bi directionnel osmoze	4	4	Coin	Plein	M	Tôlerie	ChariotAGV3.png	100.00
34	Chariot AGV frontal bi directionnel osmoze	4	4	Coin	Plein	L	Tôlerie	ChariotAGV3.png	100.00
35	Chariot AGV frontal uni directionnel C Mayor	4	4	Coin	Plein	S	Tôlerie	ChariotAGV4.png	100.00
36	Chariot AGV frontal uni directionnel C Mayor	4	4	Coin	Plein	M	Tôlerie	ChariotAGV4.png	100.00
37	Chariot AGV frontal uni directionnel C Mayor	4	4	Coin	Plein	L	Tôlerie	ChariotAGV4.png	100.00
38	Chariot AGV frontal bi directionnel C Mayor	4	4	Coin	Plein	S	Tôlerie	ChariotAGV5.png	100.00
39	Chariot AGV frontal bi directionnel C Mayor	4	4	Coin	Plein	M	Tôlerie	ChariotAGV5.png	100.00
40	Chariot AGV frontal bi directionnel C Mayor	4	4	Coin	Plein	L	Tôlerie	ChariotAGV5.png	100.00
1	Plateau roulant	1	4	Coin	Plein	S	Tôlerie	Plateauroulant.png	120.00
3	Plateau roulant	1	4	Coin	Plein	L	Tôlerie	Plateauroulant.png	80.00
\.


--
-- Name: accessoire_defaut_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accessoire_defaut_id_seq', 10, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 1, false);


--
-- Name: compatibilites_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.compatibilites_id_seq', 32, true);


--
-- Name: liste_accessoire_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.liste_accessoire_id_seq', 1, false);


--
-- Name: moyens_roulants_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.moyens_roulants_id_seq', 40, true);


--
-- Name: accessoire_defaut accessoire_defaut_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accessoire_defaut
    ADD CONSTRAINT accessoire_defaut_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: compatibilites compatibilites_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.compatibilites
    ADD CONSTRAINT compatibilites_pkey PRIMARY KEY (id);


--
-- Name: liste_accessoire liste_accessoire_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.liste_accessoire
    ADD CONSTRAINT liste_accessoire_pkey PRIMARY KEY (id);


--
-- Name: moyens_roulants moyens_roulants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.moyens_roulants
    ADD CONSTRAINT moyens_roulants_pkey PRIMARY KEY (id);


--
-- Name: accessoire_defaut accessoire_defaut_accessoire_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accessoire_defaut
    ADD CONSTRAINT accessoire_defaut_accessoire_id_fkey FOREIGN KEY (accessoire_id) REFERENCES public.liste_accessoire(id) ON DELETE CASCADE;


--
-- Name: accessoire_defaut accessoire_defaut_categorie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accessoire_defaut
    ADD CONSTRAINT accessoire_defaut_categorie_id_fkey FOREIGN KEY (categorie_id) REFERENCES public.categories(id) ON DELETE CASCADE;


--
-- Name: compatibilites compatibilites_accessoire_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.compatibilites
    ADD CONSTRAINT compatibilites_accessoire_id_fkey FOREIGN KEY (accessoire_id) REFERENCES public.liste_accessoire(id) ON DELETE CASCADE;


--
-- Name: compatibilites compatibilites_categorie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.compatibilites
    ADD CONSTRAINT compatibilites_categorie_id_fkey FOREIGN KEY (categorie_id) REFERENCES public.categories(id) ON DELETE CASCADE;


--
-- Name: moyens_roulants moyens_roulants_categorie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.moyens_roulants
    ADD CONSTRAINT moyens_roulants_categorie_id_fkey FOREIGN KEY (categorie_id) REFERENCES public.categories(id) ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

