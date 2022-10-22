from pyspark.sql.types import FloatType
from pyspark.sql import SparkSession
import pyspark.sql.functions as F
from geopy.distance import geodesic

spark = SparkSession.builder.getOrCreate()

@F.udf(returnType=FloatType())
def geodesic_udf(a_lat, a_long, b_lat, b_long):
    return geodesic((a_lat, a_long), (b_lat, b_long)).miles

df = spark.read.csv('../country-capitals.csv', header=True)
df = df.withColumn("idx", F.monotonically_increasing_id())

df = df.alias('left').join(df.alias('right'), F.col('left.idx') > F.col('right.idx')).select(
    F.col('left.CountryName'), 
    F.col('left.CapitalName'), 
    F.col('right.CountryName'), 
    F.col('right.CapitalName'),
    geodesic_udf(F.col('left.CapitalLatitude'), F.col('left.CapitalLongitude'), F.col('right.CapitalLatitude'), F.col('right.CapitalLongitude')).alias('distance')
).orderBy(F.col('distance'))
df.show(100)