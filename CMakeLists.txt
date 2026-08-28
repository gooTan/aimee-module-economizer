cmake_minimum_required(VERSION 3.16)
project(aimee_module_economizer VERSION 0.3.3 LANGUAGES NONE)

include(GNUInstallDirs)
find_program(GO_EXECUTABLE NAMES go REQUIRED)
set(MODULE_BINARY "${CMAKE_CURRENT_BINARY_DIR}/aimee-module-economizer")
add_custom_command(
    OUTPUT "${MODULE_BINARY}"
    COMMAND ${CMAKE_COMMAND} -E env CGO_ENABLED=0 ${GO_EXECUTABLE} build -trimpath
            -o "${MODULE_BINARY}" ./runtime
    WORKING_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}"
    DEPENDS
        ${CMAKE_CURRENT_SOURCE_DIR}/runtime/main.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/budget.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/cjson.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/closet.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/compact.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/condense_apply.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/condense_family.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/condense_primitives.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/condense_recognize.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/eager.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/episode.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/fold.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/gateway.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/gateway_wire.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/json_compact.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/jsonval.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/module.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/module_aux.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/module_record.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/planner_anthropic.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/planner_openai.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/proof.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/provenance.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/rail.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/recall.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/reduce.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/register.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/repair.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/modules/economizer/state.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/bus/client.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/bus/client_connect.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/bus/module_caller.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/bus/module_protocol.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/bus/module_runtime.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/bus/region.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/bus/ring.go
        ${CMAKE_CURRENT_SOURCE_DIR}/server-go/bus/wire.go
        ${CMAKE_CURRENT_SOURCE_DIR}/go.mod
        ${CMAKE_CURRENT_SOURCE_DIR}/go.sum
    VERBATIM)
add_custom_target(aimee-module-economizer ALL DEPENDS "${MODULE_BINARY}")
configure_file(grants/module.grant.in ${CMAKE_CURRENT_BINARY_DIR}/economizer.grant @ONLY)
install(PROGRAMS "${MODULE_BINARY}" DESTINATION ${CMAKE_INSTALL_BINDIR})
install(FILES ${CMAKE_CURRENT_BINARY_DIR}/economizer.grant
        DESTINATION ${CMAKE_INSTALL_DATADIR}/aimee/module-grants)
