--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.2
-- Dumped by pg_dump version 9.6.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS coopcat;
--
-- Name: coopcat; Type: DATABASE; Schema: -; Owner: coopcat_dev
--

CREATE DATABASE coopcat WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_United States.1252' LC_CTYPE = 'English_United States.1252';


ALTER DATABASE coopcat OWNER TO coopcat_dev;

\connect coopcat

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: update_num_available_positions(); Type: FUNCTION; Schema: public; Owner: coopcat_dev
--

CREATE FUNCTION update_num_available_positions() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
	BEGIN	
		UPDATE job
		SET num_available_positions = (SELECT num_available_positions - 1
						FROM job
						WHERE JOB_ID = NEW.JOBJOB_ID)
		WHERE JOB_ID = NEW.JOBJOB_ID;
		
		RETURN NEW;
	END;
$$;


ALTER FUNCTION public.update_num_available_positions() OWNER TO coopcat_dev;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: company; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE company (
    company_id numeric(19,0) NOT NULL,
    name character varying(150) NOT NULL,
    description character varying(255) NOT NULL,
    num_employees integer NOT NULL
);


ALTER TABLE company OWNER TO coopcat_dev;

--
-- Name: job; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE job (
    job_id numeric(19,0) NOT NULL,
    company_id numeric(19,0) NOT NULL,
    name character varying(100),
    rating integer,
    student_level integer,
    num_of_positions integer NOT NULL,
    deadline_date date NOT NULL,
    num_available_positions integer
);


ALTER TABLE job OWNER TO coopcat_dev;

--
-- Name: resume; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE resume (
    "- Resume_ID" numeric(19,0) NOT NULL,
    "Author- ID" numeric(19,0) NOT NULL,
    last_updated_at timestamp without time zone NOT NULL,
    "File Path" character varying(200)
);


ALTER TABLE resume OWNER TO coopcat_dev;

--
-- Name: review; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE review (
    comments character varying(300) NOT NULL,
    rating integer NOT NULL,
    works_forjobjob_id numeric(19,0) NOT NULL,
    "Works_ForUSER- ID" numeric(19,0) NOT NULL
);


ALTER TABLE review OWNER TO coopcat_dev;

--
-- Name: reviews; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE reviews (
    resume_id numeric(19,0) NOT NULL,
    moderator_id numeric(19,0) NOT NULL,
    review_date timestamp without time zone NOT NULL,
    review character varying(300) NOT NULL
);


ALTER TABLE reviews OWNER TO coopcat_dev;

--
-- Name: student; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE student (
    program character varying(200),
    school_start_date date,
    school_end_date date,
    student_level integer NOT NULL,
    user_id numeric(19,0) NOT NULL
);


ALTER TABLE student OWNER TO coopcat_dev;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: coopcat_dev
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO coopcat_dev;

--
-- Name: user; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE "user" (
    id numeric(19,0) DEFAULT nextval('user_id_seq'::regclass) NOT NULL,
    name character varying(100) NOT NULL,
    password character varying(250) NOT NULL,
    email character varying(100) NOT NULL,
    address character varying(100),
    role integer,
    CONSTRAINT "USER_email_check" CHECK (((email)::text ~* '^[A-Za-z0-9]+@[A-Za-z0-9]+.[A-Za-z0-9]+$'::text))
);


ALTER TABLE "user" OWNER TO coopcat_dev;

--
-- Name: works_for; Type: TABLE; Schema: public; Owner: coopcat_dev
--

CREATE TABLE works_for (
    jobjob_id numeric(19,0) NOT NULL,
    "USER- ID" numeric(19,0) NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL
);


ALTER TABLE works_for OWNER TO coopcat_dev;

--
-- Data for Name: company; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: job; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: resume; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: review; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: student; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--

INSERT INTO "user" (id, name, password, email, address, role) VALUES (4, 'John Doe1', 'password', 'johndoe1@gmail.com', '123 North Street', 1);
INSERT INTO "user" (id, name, password, email, address, role) VALUES (6, 'John Doe2', 'password', 'johndoe2@gmail.com', '123 North Street', 1);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: coopcat_dev
--

SELECT pg_catalog.setval('user_id_seq', 6, true);


--
-- Data for Name: works_for; Type: TABLE DATA; Schema: public; Owner: coopcat_dev
--



--
-- Name: user USER_email_key; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT "USER_email_key" UNIQUE (email);


--
-- Name: user USER_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT "USER_pkey" PRIMARY KEY (id);


--
-- Name: company company_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY company
    ADD CONSTRAINT company_pkey PRIMARY KEY (company_id);


--
-- Name: job job_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY job
    ADD CONSTRAINT job_pkey PRIMARY KEY (job_id);


--
-- Name: resume resume_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY resume
    ADD CONSTRAINT resume_pkey PRIMARY KEY ("- Resume_ID");


--
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (resume_id);


--
-- Name: works_for works_for_pkey; Type: CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY works_for
    ADD CONSTRAINT works_for_pkey PRIMARY KEY (jobjob_id, "USER- ID");


--
-- Name: works_for update_num_available_positions_trigger; Type: TRIGGER; Schema: public; Owner: coopcat_dev
--

CREATE TRIGGER update_num_available_positions_trigger AFTER INSERT ON works_for FOR EACH ROW EXECUTE PROCEDURE update_num_available_positions();


--
-- Name: job fkjob188415; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY job
    ADD CONSTRAINT fkjob188415 FOREIGN KEY (company_id) REFERENCES company(company_id);


--
-- Name: resume fkresume602495; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY resume
    ADD CONSTRAINT fkresume602495 FOREIGN KEY ("Author- ID") REFERENCES "user"(id);


--
-- Name: reviews fkreviews395856; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY reviews
    ADD CONSTRAINT fkreviews395856 FOREIGN KEY (resume_id) REFERENCES resume("- Resume_ID");


--
-- Name: reviews fkreviews794469; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY reviews
    ADD CONSTRAINT fkreviews794469 FOREIGN KEY (moderator_id) REFERENCES "user"(id);


--
-- Name: student fkstudent664699; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY student
    ADD CONSTRAINT fkstudent664699 FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE;


--
-- Name: works_for fkworks_for151853; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY works_for
    ADD CONSTRAINT fkworks_for151853 FOREIGN KEY ("USER- ID") REFERENCES "user"(id);


--
-- Name: works_for fkworks_for677765; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY works_for
    ADD CONSTRAINT fkworks_for677765 FOREIGN KEY (jobjob_id) REFERENCES job(job_id);


--
-- Name: review review; Type: FK CONSTRAINT; Schema: public; Owner: coopcat_dev
--

ALTER TABLE ONLY review
    ADD CONSTRAINT review FOREIGN KEY (works_forjobjob_id, "Works_ForUSER- ID") REFERENCES works_for(jobjob_id, "USER- ID");


--
-- PostgreSQL database dump complete
--

