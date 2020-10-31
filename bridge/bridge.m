//
//  bridge.m
//  tagger
//
//  Created by Diego Magdaleno on 10/30/20.
//

#import <Foundation/Foundation.h>
#include "bridge.h"
#include "ext/NSTaggerURL.h"
#include "ext/FileProperties.h"

#include <Foundation/Foundation.h>

@interface FileProperties : NSObject
@property (nonatomic, strong) NSString *name;
@property (nonatomic, strong) NSArray *tags;

- (id)initWithName:(NSString *)name tags:(NSArray *)tags;

@end

@implementation FileProperties
@synthesize name;
@synthesize tags;

- (id)initWithName:(NSString *)name tags:(NSArray *)tags
{
    if (self = [super init]) {
        self.name = name;
        self.tags = tags;
    }
    return self;
}
@end

typedef struct _TGFileProperties {
    NSString *name;
    NSArray *tags;
} TGFileProperties;

const TGFileProperties*
filePropertiesData(FileProperties *file){
    TGFileProperties *filePropertiesData = malloc(sizeof(TGFileProperties));
    filePropertiesData->name = file.name;
    filePropertiesData->tags = file.tags;
    return filePropertiesData;
}

const char*
NSStringToCString(NSString *s) {
    if (s == NULL) {return NULL; }

    const char *CStr = [s UTF8String];
    return CStr;
}

int
NSNumberToInt(NSNumber *i) {
    if (i == NULL ) {return 0;}
    return i.intValue;
}

unsigned long
NSArrayLen(NSArray *arrWithoutDetermined) {
    if (arrWithoutDetermined == NULL) { return 0; }

    return arrWithoutDetermined.count;
}

const void*
NSArrayItem(NSArray* arrWithDetermined, unsigned long i) {
    if (arrWithDetermined== NULL) { return NULL; }

    return [arrWithDetermined objectAtIndex:i];
}

NSArray* listFilesWithTagsAtNSURL(NSArray* path) {
    NSMutableArray *filePreliminary = [NSMutableArray array];

    for (id eachPathAsURL in path) {
        NSURL *target = [NSURL fileURLWithPath:[eachPathAsURL path]];

        if([target GetTags] != nil) {
            [filePreliminary addObject:[target path]];
        }
    } 
    return filePreliminary;
}

NSArray*
fromDictionaryToFilePropertiesArray(NSDictionary *targetDict) {
    NSMutableArray *objectList = [NSMutableArray array];
    
    
    for (id key in targetDict) {
        FileProperties *target = [[FileProperties alloc] init];
        target.name = key;
        target.tags = [targetDict objectForKey:key];
        [objectList addObject:target];
    }
    
    return objectList;
}

const NSArray*
getFilesWithCertainMacOSTag(char const*  pathRaw) {
    
    NSString *path = [NSString stringWithUTF8String:pathRaw];

    NSMutableDictionary *mappedFiles = [NSMutableDictionary dictionary];


    /*
     * Stage 1:
     * We are now on ObjC so we need to convert our Path string to
     * an NSURL
     */
    NSURL *directoryURL = [[NSURL alloc] initFileURLWithPath:path];
    
    NSArray *folders = [[NSFileManager defaultManager] contentsOfDirectoryAtURL:directoryURL includingPropertiesForKeys:NULL options:NSDirectoryEnumerationSkipsHiddenFiles error:NULL];
    
    
    for (id eachPathAsURL in folders) {
        NSURL *target = [NSURL fileURLWithPath:[eachPathAsURL path]];
        
        mappedFiles[[target path]] = [target GetTags];        
    }
    
    return fromDictionaryToFilePropertiesArray(mappedFiles);
    
}
