CREATE OR REPLACE FUNCTION "public"."avg"()
  RETURNS "pg_catalog"."trigger" AS $BODY$
BEGIN  
	IF TG_OP = 'DELETE' THEN
			INSERT INTO stock_adj(product_id)
			VALUES(OLD.product_id);
			RETURN OLD;
	END IF;
	IF TG_OP = 'UPDATE' THEN
				 IF NEW.product_id <> OLD.product_id THEN
					 INSERT INTO stock_adj(product_id)
					 VALUES(NEW.product_id);
					 INSERT INTO stock_adj(product_id)
					 VALUES(OLD.product_id);
				 ELSE
					 INSERT INTO stock_adj(product_id)
					 VALUES(NEW.product_id);
				 END IF;
		RETURN NEW;
	END IF;		
	IF TG_OP = 'INSERT' THEN
		 INSERT INTO stock_adj(product_id)
		 VALUES(NEW.product_id);
		 RETURN NEW;
	END IF;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100