package se.callistaenterprise.goblog

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import io.gatling.jdbc.Predef._
import scala.concurrent.duration._

class LoadTest extends Simulation {

    /* Place for arbitrary Scala code that is to be executed before the simulation begins. */
    before {
        println("***** My simulation is about to begin! *****")
    }

    /* Place for arbitrary Scala code that is to be executed after the simulation has ended. */
    after {
        println("***** My simulation has ended! ******")
    }

    setUp(
        Scenarios.scn_Browse.inject(rampUsers(Conf.users) over (Scenarios.rampUpTimeSecs seconds)).protocols(Conf.httpConf),
        Scenarios.scn_Create.inject(rampUsers(Conf.users) over (Scenarios.rampUpTimeSecs seconds)).protocols(Conf.httpConf),
        Scenarios.scn_JsonCreate.inject(rampUsers(Conf.users) over (Scenarios.rampUpTimeSecs seconds)).protocols(Conf.httpConf)
    )
}