package main

import "testing"

func TestPairSplitNone(t *testing.T){
    sd, s, d, _ := srcDestRead("./TestFiles/test.csv")
    if len(sd) != 0{
        t.Error("Expected 0, got ", len(sd))
    }
    if len(s) != 0{
        t.Error("Expected 0, got ", len(s))
    }
    if len(d) != 0{
        t.Error("Expected 0, got ", len(s))
    }
}

func TestPairSplitTwo(t *testing.T){
    sd, s, d, _ := srcDestRead("./TestFiles/test2.csv")
    if len(sd) != 2{
        t.Error("Expected 2, got ", len(sd))
    }
    if len(s) != 2{
        t.Error("Expected 2, got ", len(s))
    }
    if len(d) != 2{
        t.Error("Expected 2, got ", len(s))
    }
    for _, v := range s{
        if v != "hello"{
            t.Error("Array has unexpected values")
        }
    }
    for _, v := range d{
        if v != "world"{
            t.Error("Array has unexpected values")
        }
    }
    if (sd[0] != SrcDest{"hello", "world"}) || (sd[1] != SrcDest{"hello", "world"}){
        t.Error("Array has unexpected values")
    }
}

func TestFilesExistFalse(t *testing.T){
    if filesExist([]string{"./made_up_file_1", "./made_up_file_2"}){
        t.Error("Expected files to not exist")
    }
}

func TestFilesExistTrue(t *testing.T){
    if !filesExist([]string{"./TestFiles/TestDir1", "./TestFiles/TestDir2"}){
        t.Error("Expected files to exist")
    }
}

func TestUniqueFalse(t *testing.T){
    if isUnique([]string{"item1", "item2", "item1"}){
        t.Error("Expected not unique")
    }
}

func TestUniqueTrue(t *testing.T){
    if !isUnique([]string{"item1", "item2", "item3", "item4"}){
        t.Error("Expected unique")
    }
}

func TestUniqueInputFileError(t *testing.T){
    err := syncer("./TestFiles/testdup.csv")
    if err == nil {
        t.Error("Expected an error")
    }
    if err.Error() != "Destination directory has been duplicated"{
        t.Error("Expected uniqueness error")
    }
}

func TestOverlapTrue(t *testing.T){
    if !isOverlap([]string{"item1", "item2", "item3"}, []string{"item4", "item3"}){
        t.Error("Expected overlap")
    }
}

func TestOverlapFalse(t *testing.T){
    if isOverlap([]string{"item1", "item2", "item3"}, []string{"item4", "item5", "item6"}){
        t.Error("Expected no overlap")
    }
}

func TestOverlapInputFileError(t *testing.T){
    err := syncer("./TestFiles/testoverlap.csv")
    if err == nil {
        t.Error("Expected an error")
    }
    if err.Error() != "Source and destination directory list has overlap"{
        t.Error("Expected overlap error")
    }
}
