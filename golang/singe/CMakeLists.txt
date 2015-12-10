add_custom_target(build_golang_singe
        go fmt
        COMMAND go build
        COMMENT "building golang singe"
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        SOURCES singe.go
)

add_test(golang_singe_tests COMMAND go fmt && go test ${CMAKE_CURRENT_SOURCE_DIR})
